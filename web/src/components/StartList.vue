<template>
  <div>
    <div class="d-flex startlist">
      <main-table
        :headers="fields"
        :items="table"
        :fontSize="fontSize"
        v-for="(table, i) in tables"
        gridStr="minmax(30px, 1fr) minmax(200px, 6fr) minmax(80px, 1fr) minmax(150px, 2fr) minmax(110px, 2fr)"
        :key="`table${i}`"
      />
    </div>
    <v-btn
      icon
      color="indigo"
      @click="fontSize++"
    >
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn
      icon
      color="indigo"
      @click="fontSize--"
    >
      <v-icon>mdi-minus</v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import MainTable from '@/components/MainTable.vue';

@Component({
  components: { MainTable }
})
export default class StartList extends Vue {
  @Prop() private items: any;

  private fontSize: number = 12;

  get tables() {
    const items = this.items;//[...this.items, ...this.items];
    if (items.length > 50 && items.length < 100) {
      return [items.slice(0, 40), items.slice(40)];
    } else if (items.length > 100) {
      return [items.slice(0, 40), items.slice(40, 80), items.slice(80)];
    }
  }

  private fields: any = [
    {
      name: "Number",
      title: "№"
    },
    {
      name: "Name",
      title: "Фамилия имя"
    },
    {
      name: "Rating",
      title: "Рейтинг"
    },
    {
      name: "City",
      title: "Город"
    },
    {
      name: "Fide",
      title: "ID FIDE"
    },
  ];
}
</script>

<style lang="scss" scoped>
.startlist {
  gap: 30px;
}
</style>