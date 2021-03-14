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
          <h2 style="width: 960px">{{ name }}</h2>
          <h2>{{ title }} <span
              class="red--text"
              v-if="time"
            >{{ timerText }} {{ time }}</span></h2>
        </div>
      </v-col>
      <v-spacer />
      <v-col
        class="d-flex justify-end py-0"
        cols="3"
      >
        <vue-qrcode
          :value="`https://chess-results.com/Tnr${id}.aspx?lan=11`"
          :width="150"
        />
      </v-col>
    </v-row>
    <template v-if="status == 1">
      <pairs
        class="pairs"
        :pairs="pairs"
        :column="pairsColumn"
        :lineHeight="lineHeight"
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
      <finish-list :items="finishlist" :column="finishColumn" :lineHeight="lineHeight" />
    </template>
    <template v-else-if="status == 3">
      <start-list :items="startlist" />
    </template>
    <template v-else-if="status == 4">
      <div class="welcome">
        <v-carousel
          cycle
          hide-delimiter-background
          hide-delimiters
          :show-arrows="false"
          interval="10000"
          height="100%"
        >
          <v-carousel-item
            v-for="(slide, i) in slides"
            :key="i"
            reverse-transition="fade-transition"
            transition="fade-transition"
          >
            <v-sheet
              height="100%"
              color="white"
              class="black--text d-flex align-center justify-center"
            >
              <div v-html="slide" />
            </v-sheet>
          </v-carousel-item>
        </v-carousel>
      </div>
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import io from 'socket.io-client';
import Pairs from '@/components/Pairs.vue';
import Competitors from '@/components/Competitors.vue';
import StartList from '@/components/StartList.vue';
import FinishList from '@/components/FinishList.vue';
import VueQrcode from 'vue-qrcode';

enum Status {
  NotStarted = 0,
  InProgress = 1,
  Finished = 2,
  StartList = 3,
  Info = 4
}

@Component({
  components: { Pairs, Competitors, VueQrcode, StartList, FinishList },
})
export default class Home extends Vue {
  private io: any;

  private round: number = 0;
  private pairs: any = [];
  private competitors: any = [];
  private results: any = [];
  private startlist: any = [];
  private finishlist: any = [];
  private id: string = '';

  private timer: number = 0;
  private timerText: string = '';
  private timerInterval: any = null;

  private pairsColumn: string = '';
  private finishColumn: string = '';
  private lineHeight: string = '';
  private name: string = '';

  get time() {
    function pad(number: any, length: any) {
      var str = "" + number
      while (str.length < length) {
        str = '0' + str
      }
      return str
    }
    return this.timer === 0 ? null : (pad((this.timer - (this.timer % 60)) / 60, 2) + ':' + pad(this.timer % 60, 2));
  }

  private status: Status = Status.Info;
  private slides: any = [];

  get title() {
    switch (this.status) {
      case Status.StartList: return 'Стартовый лист';
      case Status.InProgress: return this.round + ' тур';
      case Status.Finished: return 'Положение после ' + this.round + ' тура';
      default: return '';
    }
  }


  async created() {
    this.io = io('http://chess-results-viewer.herokuapp.com', <any>{
      // withCredentials: false,
    });
    this.io.on('SetRound', (data: any) => {
      this.round = parseInt(data.Round);
      this.pairs = data.Pairs.Items;
      this.competitors = data.Competitors.Items;
      this.status = Status.InProgress;
    });
    this.io.on('SetStartList', (data: any) => {
      this.startlist = data.Items;
      this.status = Status.StartList;
    });
    this.io.on('SetResults', (data: any) => {
      this.finishlist = data.Competitors.Items;
      this.round = parseInt(data.Round);
      this.status = Status.Finished;
    });
    this.io.on('SetTimer', (text: any, timer: any) => {
      if (this.timerInterval) {
        clearInterval(this.timerInterval);
      }
      this.timer = parseInt(timer) * 60;
      this.timerText = text;
      this.timerInterval = setInterval(() => {
        this.timer--;
        if (this.timer === 0) {
          clearInterval(this.timerInterval);
        }
      }, 1000)
    });
    this.io.on('RemoveTimer', (data: any) => {
      if (this.time) {
        clearInterval(this.timerInterval);
        this.timer = 0;
      }
    });
    this.io.on('SetInfo', (data: any) => {
      // console.log(data);
      this.status = Status.Info;
      // this.slides = data;
      try {
        this.slides = JSON.parse(data || "['']");
      } catch (e) {

      }
    });

    let tournament: any = await Vue.$http.get('info');
    const data = JSON.parse(tournament.data.Data);
    this.pairsColumn = data.pairsColumn || '28,57';
    this.finishColumn = data.finishColumn || '37,74,150';
    this.lineHeight = data.lineHeight || '25';
    this.name = data.name || 'Кубок ГРАН–ПРИ г. Екатеринбурга по быстрым шахматам 2021 год – Этап 1';
    this.id = tournament.data.Tournament;
    switch (tournament.data.Status) {
      case 1:
        await Vue.$http.post('round', { 'Round': tournament.data.Round });
        break;
      case 2:
        await Vue.$http.post('results', { 'Round': tournament.data.Round });
        break;
      case 3:
        await Vue.$http.post('tournament', { 'Tournament': tournament.data.Tournament });
        break;
      case 4:
        await Vue.$http.post('info', { 'Info': tournament.data.Info || "['']" });
        break;
    }
  }
}
</script>

<style lang="scss" scoped>
.welcome {
  width: 100%;
  height: calc(100vh - 150px);
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