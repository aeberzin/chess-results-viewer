<template>
  <table
    class="table"
    :class="{ pair: isPair, result: isResult }"
    :style="`font-size: ${fontSize}px; line-height: ${isPair ? lineHeight + 'px' : 'normal'} !important;`"
  >
    <thead>
      <tr>
        <th
          v-for="(field, i) in headers"
          :key="i"
        >{{ field.title }}</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(item, i) in items"
        :key="i"
      >
        <td
          v-for="(field, j) in headers"
          :key="'pair' + j"
        >{{ item[field.name]}}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Prop } from 'vue-property-decorator';

@Component
export default class MainTable extends Vue {
  @Prop({ required: true }) headers!: any;
  @Prop({ required: true }) items!: any;
  @Prop({ required: false }) fontSize!: number;
  @Prop({ required: false }) isPair!: boolean;
  @Prop({ required: false }) isResult!: boolean;
  @Prop({ default: '25' }) lineHeight!: string;
}
</script>

<style lang="scss" scoped>
.table {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns:
    minmax(30px, 1fr)
    minmax(200px, 6fr)
    minmax(80px, 1fr)
    minmax(150px, 2fr)
    minmax(110px, 2fr);

  &.pair {
    grid-template-columns:
      minmax(50px, 1fr)
      minmax(150px, 6fr)
      minmax(50px, 1fr)
      minmax(50px, 1fr)
      minmax(50px, 1fr)
      minmax(150px, 6fr) !important;
  }

  &.result {
    grid-template-columns:
      minmax(50px, 1fr)
      minmax(50px, 1fr)
      minmax(150px, 6fr)
      minmax(50px, 1fr)
      minmax(50px, 1fr)
      minmax(50px, 1fr) !important;
  }
}

thead,
tbody,
tr {
  display: contents;
}
th,
td {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  align-items: center;
  display: flex;
  padding-left: 10px !important;
}
th {
  position: sticky;
  top: 0;
  background: #6c7ae0;
  text-align: left;
  font-weight: normal;
  color: white;
  position: relative;
}
td {
  color: #808080;
}

tr:nth-child(even) td {
  background: #f8f6ff;
}
</style>