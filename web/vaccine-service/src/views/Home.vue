<template>
  <div class="hero">
    <div class="tile is-child has-background-light">
      <div class="my-4">
        <p class="subtitle">Need help finding a COVID-19 vaccine in the U.S.? Call 1-800-232-0233
          (TTY 888-720-7489)</p>
      </div>
    </div>

    <div class="tile is-ancestor">
      <div class="tile is-vertical is-8">
        <div class="tile">
          <div class="tile is-parent is-vertical">
            <article class="tile is-child notification is-primary">
              <p class="title">Find a COVID-19 vaccine near you</p>
              <p class="subtitle">Use Vaccines.gov to find a location near you, then call or visit
                their website to make an appointment.</p>
              <b-button type="is-info is-light">Find COVID-19 Vaccines</b-button>
            </article>
          </div>
          <div class="tile is-parent">
            <article class="tile is-child">
              <figure class="image is-4by3">
                <img :src="require('@/assets/images/about-image.jpg')" alt="about-image.jpg">
              </figure>
            </article>
          </div>
        </div>
      </div>
      <div class="tile is-parent">
        <article class="tile is-child notification is-success">
          <div class="content">
            <p class="title">Get Help</p>
            <p class="subtitle">Visit the CDC’s Frequently Asked Questions page for information
              about COVID-19 vaccinations.</p>
            <p class="subtitle"> Get answers to questions or help finding a vaccine near you. Help
              is available in
              English, Spanish, and many other languages.</p>
            <div class="content">
              <p>1-800-232-0233 TTY 1-888-720-7489 </p>
              <p> Disability Information and Access Line (DIAL): </p>
              <p> 1-888-677-1199 DIAL@n4a.org</p>
            </div>
          </div>
        </article>
      </div>
    </div>

    <div class="tile is-parent">
      <article class="tile is-child notification is-danger">
        <p class="title">We want to make COVID-19 vaccination easy and accessible to
          everyone.</p>
        <p class="subtitle">We can do this.</p>
        <div class="content">
          <!-- Content -->
        </div>
      </article>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'Home',
  components: {},
};
</script>

<style lang="css" scoped>
.has-background-color-blue {
  background: rgb(0, 87, 183) none repeat scroll 0% 0%
}
</style>
