import Vue from 'vue';
import VueFormulate from '@braid/vue-formulate';
import Buefy from 'buefy';
import VueFormWizard from 'vue-form-wizard';
import App from './App.vue';
import store from './store';
import router from './router';
import 'buefy/dist/buefy.css';
// eslint-disable-next-line import/order
import axios from 'axios';
// eslint-disable-next-line import/order
import VueAxios from 'vue-axios';
import 'vue-form-wizard/dist/vue-form-wizard.min.css';

Vue.use(VueFormWizard);
Vue.use(Buefy);
Vue.use(VueAxios, axios);
Vue.use(VueFormulate);

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');
