<template>
  <div id="app">
    <div v-for="(pair,i) in pairs" :key="i"><b>{{ pair.Table }}.</b> {{ pair.WName }} - {{ pair.BName }} | {{ pair.Result }}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import HelloWorld from './components/HelloWorld.vue';

@Component({
  components: {
    HelloWorld,
  },
})
export default class App extends Vue {
  private pairs: any = [];

  created() {
    fetch("http://localhost:3000/api/pairs")
      .then(response => response.json())
      .then(data => {
        this.pairs = data;
      });
  }
}
</script>

<style lang="scss">
body {
  margin: 0;
  font-family: BlinkMacSystemFont, -apple-system, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", "Helvetica", "Arial", sans-serif;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100vh;
  display: grid;
  & > div {
    border: 1px solid black;
    &:nth-child(2n) {
      background: #f0f0f0;
    }
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>
