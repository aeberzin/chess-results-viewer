import Vue from 'vue';
import axios from 'axios';

const instance = axios.create({
  baseURL: `${window.location.protocol}//${window.location.hostname}/api`,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// instance.interceptors.request.use(
//   (config: any) => {
//     const token: any = store.state.auth.token;
//     if (token) {
//       config.headers.Authorization = `TOKEN ${token}`;
//     }
//     return config;
//   },
//   (error: any) => Promise.reject(error),
// );

const axiosPlugin = {
  install(vueObj: any) {
    vueObj.$http = instance;
    Object.defineProperty(Vue.prototype, '$http', { value: instance });
  },
};

Vue.use(axiosPlugin);


declare module 'vue/types/vue' {
  export interface VueConstructor {
    $http: typeof instance;
  }
  export interface Vue {
    $http: typeof instance;
  }
}