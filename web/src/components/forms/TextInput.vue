<script>
export default {
  name: "TextInput",
  props: {
    name: String,
    type: String,
    label: String,
    placeholder: String,
    required: String,
    min: String,
    max: String,
    value: String,
  },
}
</script>

<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{label}}</label>
    <input :type="type" :name="name" :placeholder="placeholder"
           :value="value" :min="min" :max="max"
           :autocomplete="name + '-new'" :required="required"
           class="form-control">
  </div>
</template>

<style scoped>

</style>