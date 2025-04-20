<script>
import security from "@/components/security";
export default {
  name: "BooksAdmin",
  beforeMount() {
    security.requireToken()
  },
}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Books Admin</h1>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>