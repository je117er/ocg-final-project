import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const ID_TOKEN_KEY = 'id_token';

export default new Vuex.Store({
  state: {
    isAuthenticated: false,
    errors: null,
    user: {},
  },
  mutations: {
    setUser(state, user) {
      console.log(user);
      state.user = user;
      state.isAuthenticated = true;
      window.localStorage.setItem(ID_TOKEN_KEY, user.Token);
    },
    clearAuth(state) {
      state.user = {};
      state.isAuthenticated = {};
      window.localStorage.removeItem(ID_TOKEN_KEY);
    },
  },
  actions: {
    async login({ commit }, credentials) {
      try {
        const response = await Vue.axios.post('http://localhost:8088/login', credentials);
        const user = response.data;
        commit('setUser', user);
        return true;
      } catch (err) {
        console.log(err);
        return false;
      }
    },
    async logout({ commit }) {
      await window.localStorage.removeItem(ID_TOKEN_KEY);
      commit('clearAuth');
      return true;
    },
  },
  modules: {},
});
