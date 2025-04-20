<script>
import security from "@/components/security";
export default {
  name: "BookEdit",
  beforeMount() {
    security.requireToken()
  }
}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Book Edit</h1>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>