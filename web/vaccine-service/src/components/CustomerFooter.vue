<template>
  <footer class="footer has-background-color-dark ">
    <div class="content has-text-centered">
      <div class="contact">
        <ul>
          <li><strong>Contact Us</strong></li>
          <li><strong>COVID-19 Vaccine FAQ</strong></li>
          <li><strong>Provider Resources </strong></li>
        </ul>
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
ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  display: inline-flex;
}

li {
  padding: 16px;
}
</style>
