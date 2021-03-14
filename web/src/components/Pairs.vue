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
        :lineHeight="lineHeight"
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
  @Prop({ required: false, default: '' }) private column: any;
  @Prop({ required: false, default: '25' }) private lineHeight: any;

  private fontSize: number = 16;

  get tables() {
    let current = 0;
    return this.column.split(',').map((column: string) => {
      this.pairs.slice(current, parseInt(column));
      current = parseInt(column);
      console.log(current, column);
    });
    // return [this.pairs.slice(0, 28), this.pairs.slice(28, 57)]
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
      name: "WRes",
      title: "Очки"
    },
    {
      name: "Result",
      title: ' '
    },
    {
      name: "BRes",
      title: "Очки"
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