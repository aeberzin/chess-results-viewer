<template>
  <v-container>
    <h2 class="mb-2">Текущий статус: {{ status }}</h2>
    <v-text-field
      v-model="tournament"
      label="Турнир"
      :placeholder="currentTournament"
    >
      <template #append-outer>
        <button @click="setTournament">Установить</button>
      </template>
    </v-text-field>
    <v-text-field
      v-model="round"
      label="Текущий тур"
      :placeholder="currentRound"
    >
      <template #append-outer>
        <button @click="setRound">Установить</button>
      </template>
    </v-text-field>
    <v-text-field
      v-model="finishedRound"
      label="Результаты после тура"
      :placeholder="finishedRound"
    >
      <template #append-outer>
        <button @click="setResults">Установить</button>
      </template>
    </v-text-field>
    <v-btn
      class="mb-6"
      @click="setStart"
    >Установить стартовый лист</v-btn>
    <h3>Таймер</h3>
    <v-text-field
      v-model="timer.text"
      label="Текст"
    />
    <v-text-field
      v-model="timer.time"
      label="Время (мин.)"
    />
    <v-btn
      class="mb-6"
      @click="setTimer"
    >Установить</v-btn>
    <v-btn
      class="mb-6 ml-2"
      @click="removeTimer"
    >Удалить</v-btn>
    <h3 class="mb-2">Объявления</h3>
    <editor
      v-model="infoItems[i]"
      v-for="(info,i) in infoItems"
      :key="`editor${i}`"
    />
    <v-btn
      class="my-6"
      @click="infoItems.push('')"
    >Добавить</v-btn>
    <v-btn
      class="my-6 ml-2"
      @click="infoItems.pop()"
    >Удалить</v-btn>
    <v-btn
      class="my-6 ml-2"
      @click="setInfo"
    >Установить</v-btn>
    <v-text-field
      v-model="name"
      label="Название"
    >
      <template #append-outer>
        <button @click="setData">Установить</button>
      </template>
    </v-text-field>
    <v-text-field
      v-model="pairsColumn"
      label="Колонки пар"
    >
      <template #append-outer>
        <button @click="setData">Установить</button>
      </template>
    </v-text-field>
    <v-text-field
      v-model="finishColumn"
      label="Колонки положения"
    >
      <template #append-outer>
        <button @click="setData">Установить</button>
      </template>
    </v-text-field>
    <v-text-field
      v-model="lineHeight"
      label="Ширина строки"
    >
      <template #append-outer>
        <button @click="setData">Установить</button>
      </template>
    </v-text-field>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Editor from '@/components/Editor.vue';

@Component({
  components: { Editor }
})
export default class Dashboard extends Vue {
  private tournament: string = "";
  private round: string = "";
  private finishedRound: string = "";
  private status: string = "Неопределен";

  private currentTournament: string = "";
  private currentRound: string = "";

  private infoItems: any = [''];

  private pairsColumn: string = '';
  private finishColumn: string = '';
  private lineHeight: string = '';

  private name: string = '';

  private timer: any = {
    text: '',
    time: '',
  }

  async setTimer() {
    await Vue.$http.post('timer', { 'Text': this.timer.text, 'Time': this.timer.time });
  }

  async removeTimer() {
    await Vue.$http.get('timer');
  }

  async setInfo() {
    await Vue.$http.post('info', { 'Info': JSON.stringify(this.infoItems) });
    this.status = "Информация";
  }

  async setData() {
    await Vue.$http.post('data', {      'Info': JSON.stringify({
        pairsColumn: this.pairsColumn,
        finishColumn: this.finishColumn,
        lineHeight: this.lineHeight,
        name: this.name,
      })
    });
    // this.status = "Информация";
  }

  async setRound() {
    await Vue.$http.post('round', { 'Round': this.round });
    this.currentRound = this.round;
    this.round = this.round;
    this.status = "Тур " + this.round;
  }

  async setStart() {
    await Vue.$http.post('tournament', { 'Tournament': this.currentTournament });
    this.status = "Стартовый лист";
  }

  async setTournament() {
    await Vue.$http.post('tournament', { 'Tournament': this.tournament });
    this.currentTournament = this.tournament;
    this.tournament = "";
    this.status = "Стартовый лист";
  }

  async setResults() {
    await Vue.$http.post('result', { 'Round': this.finishedRound });
    this.currentTournament = this.tournament;
    this.tournament = "";
    this.status = "Результаты после " + this.finishedRound + " тура";
  }

  async created() {
    let tournament: any = await Vue.$http.get('info');
    this.currentTournament = tournament.data.Tournament;
    this.currentRound = tournament.data.Round;
    this.infoItems = JSON.parse(tournament.data.Info || "['']");
    const data = JSON.parse(tournament.data.Data || "{}");
    this.pairsColumn = data.pairsColumn || '';
    this.lineHeight = data.lineHeight || '';
    this.finishColumn = data.finishColumn || '';
    this.name = data.name || '';
    switch (tournament.data.Status) {
      case 1: this.status = "Тур " + this.round; break;
      case 2: this.status = "Результаты после " + this.round + " тура"; break;
      case 3: this.status = "Стартовый лист"; break;
      case 4: this.status = "Информация"; break;
      default: this.status = "Неопределен";
    }
  }
}
</script>