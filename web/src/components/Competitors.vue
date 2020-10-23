<template>
    <table class="competitors">
      <thead>
        <tr>
          <th v-for="(field, i) in fields" :key="i">
            {{ field.title }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(competitor, i) in competitors" :key="i">
          <td v-for="(field, j) in fields" :key="'pair' + j"> {{ competitor[field.name]}} </td>
        </tr>
      </tbody>
    </table>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class Competitors extends Vue {
    @Prop() private competitors: any;
    private fields: any = [
        {
        name: "Finish",
        title: "Место"
        },
        {
        name: "Start",
        title: "№"
        },
        {
        name: "Name",
        title: "Игрок"
        }
    ];
}
</script>

<style lang="scss" scoped>
.competitors {
  width: 100%;
  display: grid;
  grid-template-columns: 
    minmax(70px, 70px)
    minmax(70px, 70px)
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