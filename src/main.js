import Vue from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import App from './App.vue';
import router from './router';
import store from './store';

import './assets/scss/index.scss';

Vue.config.productionTip = false;

Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

// ESLint可以避免语法错误
// const name = 'myzhang';

// function sayHello(who) {
//   console.log(`hello ${who}`);
// }
// sayHello(name);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
