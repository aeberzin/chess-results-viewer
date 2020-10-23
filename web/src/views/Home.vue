<template>
  <div>
    <template v-if="status == 1">
      <div class="progress">
        <h2 class="round">Тур {{ round }}</h2>
        <pairs class="pairs" :pairs="pairs" />
        <h2 class="competition">Положение после {{ round - 1 }} тура</h2>
        <competitors class="competitors" :competitors="competitors" />
      </div>
    </template>
    <template v-else-if="status == 0">
      <div class="welcome">
        <h1>Добро пожаловать!</h1>
      </div>
    </template>
    <template v-else-if="status == 2">

    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import io from 'socket.io-client';
import Pairs from '@/components/Pairs.vue';
import Competitors from '@/components/Competitors.vue';

enum Status {
  NotStarted = 0,
  InProgress = 1,
  Finished = 2,
}

@Component({
  components: { Pairs, Competitors },
})
export default class Home extends Vue {
  private io: any;

  private round: number = 0;
  private pairs: any = [];
  private competitors: any = [];
  private results: any = [];
  private startlist: any = [];

  private status: Status = Status.NotStarted;


  created() {
    this.io = io(window.location.hostname+':3000');
    this.io.on('SetRound', (data) => {
      this.round = parseInt(data.Round);
      this.pairs = data.Pairs.Items;
      this.competitors = data.Competitors.Items;
      this.status = Status.InProgress;
    });
    this.io.on('SetRound', (data) => {
      this.round = parseInt(data.Round);
      this.results = data.Players.Items;
      this.status = Status.Finished;
    });
  }
}
</script>

<style lang="scss" scoped>
.welcome {
  width: 100%;
  height: 100vh;
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