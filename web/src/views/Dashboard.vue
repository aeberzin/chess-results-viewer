<template>
    <div>
        <h2>Текущий статус: {{ status }}</h2>
        <div>
            <label>Турнир</label>
            <input v-model="tournament" type="text" :placeholder="currentTournament">
            <button @click="setTournament">Установить</button>
        </div>
        <div>
            <label>Текущий тур</label>
            <input v-model="round" type="text" :placeholder="currentRound">
            <button @click="setRound">Установить</button>
        </div>
        <div>
            <label>Результаты после тура</label>
            <input v-model="finishedRound" type="text">
            <button @click="setResults">Установить</button>
        </div>
        <div>
            <button @click="setStart">Установить стартовый лист</button>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class Dashboard extends Vue {
    private tournament: string = "";
    private round: string = "";
    private finishedRound: string = "";
    private status: string = "Неопределен";

    private currentTournament: string = "";
    private currentRound: string = "";

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
        await Vue.$http.post('results', { 'Round': this.finishedRound });
        this.currentTournament = this.tournament;
        this.tournament = "";
        this.status = "Результаты после " + this.finishedRound + " тура";
    }

    async created() {
        let tournament: any = await Vue.$http.get('info');
        this.currentTournament = tournament.data.Tournament;
        this.currentRound = tournament.data.Round;
    }
}
</script>