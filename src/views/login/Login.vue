<!-- eslint-disable keyword-spacing -->
<!-- eslint-disable vuejs-accessibility/label-has-for -->
<!-- eslint-disable max-len -->
<!-- eslint-disable vue/multi-word-component-names -->
<template>
    <div class="login">
        <form>
            <div class="form-group">
                <label for="user.tele">Phone Number</label>
                <input type="number" class="form-control">
            </div>
            <div class="form-group">
                <label for="user.password">Password</label>
                <input type="password" class="form-control">
            </div>
            <div class="form-group form-check">
                <input type="checkbox" class="form-check-input" id="exampleCheck1">
                <label class="form-check-label" for="exampleCheck1">Check me out</label>
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
        </form>
    </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        name: '',
        tele: '',
        password: '',
      },
    };
  },
};
</script>

<style lang="scss" scoped>

</style>
