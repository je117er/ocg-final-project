<template>
  <header>
    <b-navbar v-bind:fixed-top="true">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img :src="require('@/assets/images/logo.svg')" alt="logo.svg">
          <div class="title is-4">Vaccine-x</div>
        </b-navbar-item>
      </template>
      <template #start class="text-right" v-if="!isAuthenticated">
        <b-navbar-item
          tag="router-link"
          :active="true"
          :to="{ name: 'Home'}">
          Home
        </b-navbar-item>
        <b-navbar-dropdown label="Vaccines" v-bind:hoverable="true">

          <b-navbar-item
            tag="router-link"
            :active="true"
            :to="{ name: 'product-detail', params: {id: vaccine.ID } }"
            v-for="vaccine in vaccines"
            :key="vaccine.ID"
          >
            {{ vaccine.Slug }}
          </b-navbar-item>

        </b-navbar-dropdown>
        <b-navbar-item
          tag="router-link"
          :active="true"
          :to="{ name: 'Register' }"
        >
          Register
        </b-navbar-item>
        <b-navbar-item href="#">
          Search
        </b-navbar-item>
        <b-navbar-item tag="router-link" :active="true" :to="{name: 'login'}">
          Login
        </b-navbar-item>
      </template>
      <template #start class="text-right" v-else>
        <b-navbar-item
          tag="router-link"
          :active="true"
          :to="{ name: 'Home'}">
          Home
        </b-navbar-item>
        <b-navbar-item tag="router-link" :active="true" :to="{name: 'customer-home'}">
          Customer
        </b-navbar-item>
        <b-navbar-item tag="router-link" :active="true" :to="{name: 'clinic-home'}">
          Clinic
        </b-navbar-item>
        <b-navbar-item :active="true" @click="logout">
          Logout
        </b-navbar-item>
      </template>
    </b-navbar>
  </header>
</template>

<script>
import Vue from 'vue';

export default {
  name: 'CustomerHeader',
  data() {
    return {
      vaccines: null,
    };
  },
  methods: {
    async logout() {
      await this.$store.dispatch('logout');
    },
  },
  created() {
    Vue.axios.get('http://localhost:8088/products').then((response) => {
      this.vaccines = response.data;
    }).catch((err) => {
      console.log(err);
    });
  },
  computed: {
    isAuthenticated() {
      return this.$store.state.isAuthenticated;
    },
  },
};
</script>

<style scoped>

</style>
