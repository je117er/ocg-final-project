import Vue from 'vue';
import VueFormulate from '@braid/vue-formulate';
import Buefy from 'buefy';
import axios from 'axios';
import VueAxios from 'vue-axios';
import VueFormWizard from 'vue-form-wizard';
import 'vue-form-wizard/dist/vue-form-wizard.min.css';
import App from './App.vue';
import store from './store';
import router from './router';
import 'buefy/dist/buefy.css';

Vue.use(VueFormWizard);
Vue.use(Buefy);
Vue.use(VueAxios, axios);
Vue.use(VueFormulate);

Vue.config.productionTip = false;

router.beforeEach(async (to, from, next) => {
  axios.defaults.headers.common['x-access-token'] = window.localStorage.getItem('id_token');
  next();
});

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');
