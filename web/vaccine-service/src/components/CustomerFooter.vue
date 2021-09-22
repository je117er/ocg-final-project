<template>
  <footer class="footer has-background-color-dark ">
    <div class="content has-text-centered">
      <div class="contact">
        <span><strong>Contact Us</strong></span>
        <span><strong>COVID-19 Vaccine FAQ</strong></span>
        <span><strong>Provider Resources </strong></span>
      </div>
      <p><strong>copyright @ 2021</strong></p>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'CustomerFooter',
};
</script>

<style scoped>
.has-background-color-dark {
  background: rgb(6, 34, 53);
}

.contact {
  margin: 25px auto;
}

strong {
  color: rgb(255, 255, 255);
  margin: 25px auto;
}
</style>
