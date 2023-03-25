<template>
  <div id="app">
    <navbar />
    <router-view/>
  </div>
</template>
<script>
import Navbar from './views/layout/Navbar.vue';

export default {
  components: { Navbar },
  data() {

  },
};
</script>
<style lang="scss">

</style>
