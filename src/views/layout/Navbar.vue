<!-- eslint-disable max-len -->
<!-- eslint-disable vue/multi-word-component-names -->
<template>
    <!-- As a link -->
  <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
    <a class="navbar-brand" href="#">
      <img src="../../images/golang.png" width="30" height="30" alt="">Gin&&Vue
    </a>
    <div class="form-inline" id="navbarNav">
    <ul class="navbar-nav">
      <li class="nav-item active">
        <a class="nav-link" href="#" @click="$router.replace({name:'login'})">Login<span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#" @click="$router.replace({name:'register'})">Register<span class="sr-only">(current)</span></a>
      </li>
    </ul>
  </div>
  </nav>
</template>

<script>
export default {

};
</script>

<style lang="scss" scoped>

</style>
