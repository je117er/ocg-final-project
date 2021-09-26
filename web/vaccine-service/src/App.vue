<template>
  <div id="app">
    <CustomerHeader id="header"/>
    <div id="content">
      <router-view/>
    </div>
    <CustomerFooter id="footer"/>
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#app {
  margin: 0;
  display: flex;
  flex-direction: column;
  min-height: calc(100vh - 60px);
}

#content {
  flex: 1;
}

#footer {
  height: 60px;
}
</style>
<script>
import CustomerHeader from '@/components/CustomerHeader.vue';
import CustomerFooter from '@/components/CustomerFooter.vue';

export default {
  components: { CustomerHeader, CustomerFooter },
};
</script>
