<template>
  <div>
    <div class="d-flex main">
      <main-table
        :headers="fields"
        :items="table"
        :fontSize="fontSize"
        v-for="(table, i) in tables"
        :isPair="true"
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
  components: { MainTable },
})
export default class Pairs extends Vue {
  @Prop() private pairs: any;

  private fontSize: number = 16;

  get tables() {
    return [this.pairs.slice(0, 25), this.pairs.slice(25, 50)]
  }

  private fields: any = [
    {
      name: "Table",
      title: "№"
    },
    {
      name: "WName",
      title: "Белые"
    },
    {
      name: "BName",
      title: "Черные"
    }
  ];
}
</script>

<style lang="scss" scoped>
.main {
  gap: 30px;
}
</style>