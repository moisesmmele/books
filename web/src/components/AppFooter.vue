<script>
export default {
  name: "AppFooter"
}
</script>

<template>
  <footer class="d-flex align-items-center justify-content-center mt-5">Copyright &copy; {{new Date().getFullYear()}}</footer>
</template>

<style scoped>
  footer {
    background: rgba(33,37,41,1);
    color: white;
    min-height: 3em;
    font-size: 0.8em;
  }
</style>