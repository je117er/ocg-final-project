<template>
  <header>
    <b-navbar v-bind:fixed-top="true">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img :src="require('@/assets/images/logo.svg')" alt="logo.svg">
          <div class="title is-4">Vaccine-x</div>
        </b-navbar-item>
      </template>
      <template #start class="text-right">
        <b-navbar-item href="#">
          Home
        </b-navbar-item>
        <b-navbar-dropdown label="Vaccines" v-bind:hoverable="true">

          <b-navbar-item v-for="vaccine in vaccines" :key="vaccine.ID">
<!--            <template v-if="vaccine.ID">-->
<!--              <router-link :to="{name: 'product-detail', query: {id: vaccine.ID}}">-->
                <router-link :to="{name: 'product-detail', params: {id: vaccine.ID}}">
<!--                {{ vaccine.Slug }}-->
                  {{ vaccine.Slug }}
              </router-link>
<!--            </template>-->
          </b-navbar-item>

        </b-navbar-dropdown>
        <b-navbar-item href="#">
          Register
        </b-navbar-item>
        <b-navbar-item href="#">
          Search
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
  methods: {},
  created() {
    Vue.axios.get('http://localhost:8088/products').then((response) => {
      this.vaccines = response.data;
    }).catch((err) => {
      console.log(err);
    });
  },
};
</script>

<style scoped>

</style>
