<template>
  <table class="pairs">
    <thead>
      <tr>
        <th v-for="(field, i) in fields" :key="i">{{ field.title }}</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(pair, i) in pairs" :key="i">
        <td v-for="(field, j) in fields" :key="'pair' + j">{{ pair[field.name]}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class Pairs extends Vue {
  @Prop() private pairs: any;
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
.pairs {
  width: 100%;
  display: grid;
  grid-template-columns:
    minmax(70px, 1fr)
    minmax(150px, 6fr)
    minmax(150px, 6fr);
}

thead,
tbody,
tr {
  display: contents;
}
th,
td {
  padding: 10px;
  line-height: 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  align-items: center;
  display: flex;
}
th {
  position: sticky;
  top: 0;
  background: #6c7ae0;
  text-align: left;
  font-weight: normal;
  font-size: 1.1rem;
  color: white;
  position: relative;
}
td {
  padding-top: 10px;
  padding-bottom: 10px;
  color: #808080;
}

tr:nth-child(even) td {
  background: #f8f6ff;
}
</style>