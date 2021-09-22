import Vue from 'vue';
import Buefy from 'buefy';
import App from './App.vue';
import store from './store';
import router from './router';
import 'buefy/dist/buefy.css';
// eslint-disable-next-line import/order
import axios from 'axios';
// eslint-disable-next-line import/order
import VueAxios from 'vue-axios';

Vue.use(Buefy);
Vue.use(VueAxios, axios);

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');
