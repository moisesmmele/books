<script>
export default {
  name: "FormTag",
  props: {
    name: String,
    action: String,
    method: String,
    event: String,
  },
  methods: {
    submit() {
      let myForm = this.$refs[this.$props.name];
      if (myForm.checkValidity()) {
        this.$emit(this.$props.event);
      }
      myForm.classList.add("was-validated");
    }
  },
  mounted() {
    console.log(this);
  }
}
</script>

<template>
  <form @submit.prevent="submit" :event="event" :ref="name" :action="action" :method="method" autocomplete="off" class="needs-validation" novalidate>
    <slot/>
  </form>
</template>

<style scoped>

</style>