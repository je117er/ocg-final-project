<template>
  <div>
    <div class="container is-fullhd">
      <h1 v-if="loading">Loading...</h1>
      <div v-else class="notification">
        <div class="tile is-ancestor">
          <div class="tile is-vertical">
            <div class="tile">
              <div class="tile is-parent">
                <article class="tile is-child notification is-info">
                  <figure class="image is-4by3">
                    <img :src="require('@/assets/images/logo.svg')" alt="logo.svg">
                  </figure>
                </article>
              </div>
            </div>
            <div class="tile is-parent">
              <article class="tile is-child notification is-danger">
                <p class="title">{{ productDetail.Slug }}</p>
                <p class="subtitle">{{ productDetail.Name }}</p>
                <div class="content has-text-left">
                  <p><strong>Price:</strong> {{ productDetail.Price }} $</p>
                  <p><strong>Vender: </strong> {{ productDetail.Vendor }}</p>
                  <p><strong>LotNumber: </strong>{{ productDetail.LotNumber }}</p>
                  <p><strong>ExpiryDate: </strong>{{ productDetail.ExpiryDate }}</p>
                </div>
              </article>
            </div>
          </div>
          <div class="tile is-parent  is-8">
            <article class="tile is-child notification has-text-left">
              <div class="content">
                <div class="content">
                  <p><strong>VaccineType: </strong>{{ productDetail.VaccineType }}</p>
                  <p><strong>Background: </strong>{{ productDetail.Background }}</p>
                  <p><strong>RegulatoryActions: </strong>{{ productDetail.RegulatoryActions }}</p>
                  <p><strong>SafetyStatus: </strong>{{ productDetail.SafetyStatus }}</p>
                </div>
              </div>
            </article>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';

export default {
  name: 'ProductDetail.vue',
  data() {
    return {
      productDetail: {},
      loading: false,
      // deep: true,
    };
  },
  created() {
    this.loading = true;
    this.fetchData();
  },
  watch: {
    $route: 'fetchData',
  },
  methods: {
    async fetchData() {
      const fetchedID = this.$route.params.id;
      this.productDetail = null;
      try {
        const response = await Vue.axios.get(`http://localhost:8088/product/${this.$route.params.id}`);
        if (this.$route.params.id !== fetchedID) return;
        console.log(response);
        this.productDetail = response.data;
        this.loading = false;
      } catch (e) {
        console.log(e);
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>

</style>
