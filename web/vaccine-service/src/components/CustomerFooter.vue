<template>
  <footer class="footer">
    <div class="contact">
      <ul>
        <li><strong>Contact Us</strong></li>
        <li><strong>COVID-19 Vaccine FAQ</strong></li>
        <li><strong>Provider Resources </strong></li>
      </ul>
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

.footer {
  background: rgb(6, 34, 53);
}

.contact strong {
  color: rgb(255, 255, 255);
  margin: 10px 0;

}

.contact li {
  list-style-type: none;
  margin: 0;
  overflow: hidden;
  display: inline-flex;
  padding: 0 16px;
}
</style>
