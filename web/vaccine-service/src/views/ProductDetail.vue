<template>
  <div>
    <div class="container is-fullhd">
      <template v-if="!loading">Loading...</template>
      <template v-else>
        <div class="notification">
          <div class="tile is-ancestor">
            <div class="tile is-vertical">
              <div class="tile">
                <div class="tile is-parent">
                  <article class="tile is-child is-info">
                    <figure class="image is-4by3">
<!--                      <img :src="require('@/assets/images/logo.svg')" alt="logo.svg">-->
<!--                      <img :src="require('@/assets/images/pfizer.jpg')" alt="logo.svg">-->
                      <img :src="productDetail.Image" alt="image" />
<!--                          <img src="https://www.google.com/url?sa=i&url=https%3A%2F%2Fbaotintuc.vn%2Fthe-gioi%2Fpfizer-phat-trien-vaccine-covid19-dac-tri-bien-the-delta-20210825102753231.htm&psig=AOvVaw0OgY-aeaRP8I-OhT7q2Syy&ust=1632387632482000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCMi6mv2bkvMCFQAAAAAdAAAAABAD" alt="image"/>-->
                    </figure>
                  </article>
                </div>
              </div>
              <div class="tile is-parent">
                <article class="tile is-child notification is-danger">
                  <p class="title">{{ productDetail.Name }}</p>
                  <p class="subtitle">{{ productDetail.AlternateName }}</p>
                  <div class="content has-text-left">
                    <p><strong>Vaccine Type: </strong>{{ productDetail.VaccineType }}</p>
                    <p><strong>Antigen Nature: </strong>{{ productDetail.AntigenNature }}</p>
                    <p><strong>Price:</strong> ${{ productDetail.Price }}</p>
                    <p><strong>Vendor: </strong> {{ productDetail.Vendor }}</p>
                    <p><strong>Dose: </strong>{{productDetail.Dose}}</p>
                    <p><strong>Diluent: </strong>{{productDetail.Diluent}}</p>
                    <p><strong>Immunization schedule: </strong>
                      {{productDetail.ImmunizationSchedule}} doses</p>
                    <p><strong>Adjuvant: </strong>{{productDetail.Adjuvant}}</p>
                    <p><strong>Routes of administration: </strong>
                      {{productDetail.RouteOfAdministration}}</p>
                    <p><strong>Minimum Interval: </strong>
                    {{productDetail.MinimumInterval}} days</p>
                    <p><strong>Extended Interval: </strong>
                      {{productDetail.ExtendedInterval}} days</p>
                    <p><strong>Authorized Interval: </strong>
                      {{productDetail.AuthorizedInterval}} days</p>
                    <p><strong>Storage requirements: </strong>
                      {{productDetail.StorageRequirements}}</p>
                  </div>
                </article>
              </div>
            </div>
            <div class="tile is-parent  is-8">
              <article class="tile is-child notification has-text-left">
                <div class="content">
                  <div class="content">
                    <p><strong>Background: </strong>{{ productDetail.Background }}</p>
<!--                    <p><strong>Regulatory Actions: </strong>-->
<!--                      {{ productDetail.RegulatoryActions-->
<!--                        .slice(0, productDetail.RegulatoryActions.length-2) }}</p>-->
                    <p><strong>Safety Status: </strong>{{ productDetail.SafetyStatus }}</p>
                    <p><strong>Distribution: </strong>
                      {{ productDetail.Distribution}}</p>
                    <p><strong>Funding: </strong>{{ productDetail.Funding }}</p>
                  </div>
                </div>
              </article>
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';

export default {
  name: 'ProductDetail.vue',
  data() {
    return {
      productDetail: null,
      loading: false,
    };
  },
  computed: {
  },
  created() {
    this.fetchData();
  },
  watch: {
    $route: 'fetchData',
  },
  methods: {
    async fetchData() {
      const fetchedID = this.$route.params.id;
      try {
        const response = await Vue.axios.get(`http://localhost:8088/product/${this.$route.params.id}`);
        if (this.$route.params.id !== fetchedID) return;

        console.log(response.data);
        this.productDetail = response.data;
        this.loading = true;
      } catch (e) {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>

</style>
