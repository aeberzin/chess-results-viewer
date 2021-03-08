<template>
  <div>
    <v-row class="d-flex px-10 align-center">
      <v-col
        cols="9"
        class="d-flex align-center"
      >
        <img
          :src="require('@/assets/fshso.jpg')"
          width="120"
        />
        <div class="ml-6">
          <h2>Кубок ГРАН–ПРИ г. Екатеринбурга по быстрым шахматам 2021 год – Этап 1</h2>
          <h2>{{ title }}</h2>
        </div>
      </v-col>
      <v-spacer />
      <v-col
        class="d-flex justify-end py-0"
        cols="3"
      >
        <vue-qrcode
          :value="url"
          :width="150"
        />
      </v-col>
    </v-row>
    <template v-if="status == 1">
      <pairs
        class="pairs"
        :pairs="pairs"
      />
      <!-- <h2 class="competition">Положение после {{ round - 1 }} тура</h2> -->
      <!-- <competitors class="competitors" :competitors="competitors" /> -->
    </template>
    <template v-else-if="status == 0">
      <div class="welcome">
        <h1>Добро пожаловать!</h1>
      </div>
    </template>
    <template v-else-if="status == 2">

    </template>
    <template v-else-if="status == 3">
      <start-list :items="startlist" />
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import io from 'socket.io-client';
import Pairs from '@/components/Pairs.vue';
import Competitors from '@/components/Competitors.vue';
import StartList from '@/components/StartList.vue';
import VueQrcode from 'vue-qrcode';

enum Status {
  NotStarted = 0,
  InProgress = 1,
  Finished = 2,
  StartList = 3,
}

@Component({
  components: { Pairs, Competitors, VueQrcode, StartList },
})
export default class Home extends Vue {
  private io: any;

  private round: number = 0;
  private pairs: any = [];
  private competitors: any = [];
  private results: any = [];
  private startlist: any = [];

  private url: string = 'https://chess-results.com/Tnr551049.aspx?lan=11'

  private status: Status = Status.NotStarted;

  get title() {
    switch (this.status) {
      case Status.StartList: return 'Стартовый лист';
      case Status.InProgress: return this.round + ' тур';
      case Status.Finished: return 'Положение после ' + this.round + ' тура';
      default: return 'Новая страница';
    }
  }


  created() {
    this.io = io('localhost' + ':3000', {
      withCredentials: false
    });
    this.io.on('SetRound', (data: any) => {
      this.round = parseInt(data.Round);
      this.pairs = data.Pairs.Items;
      this.competitors = data.Competitors.Items;
      this.status = Status.InProgress;
    });
    this.io.on('SetStartList', (data: any) => {
      console.log(data);
      this.startlist = data.Items;
      this.status = Status.StartList;
      // this.round = parseInt(data.Round);
      // this.results = data.Players.Items;
      // this.status = Status.Finished;
    });
  }
}
</script>

<style lang="scss" scoped>
.welcome {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.round {
  grid-area: round;
}
.competition {
  grid-area: competition;
  // @media (max-width: 1024px) {
  //   display: none;
  // }
}
.pairs {
  grid-area: pairs;
}
.list {
  grid-area: list;
  @media (max-width: 1024px) {
    display: none;
  }
}
.progress {
  margin: 15px;
  @media (min-width: 1024px) {
    display: grid;
    height: calc(100vh - 30px);
    grid-template:
      "round competition" 50px
      "pairs list" minmax(200px, auto) / 1fr 1fr;
    grid-gap: 10px;
  }
}
</style>