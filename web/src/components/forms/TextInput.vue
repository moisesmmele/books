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
    modelValue: String,
    help: String,
  },
}
</script>

<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{label}}</label>
    <input :type="type" :name="name" :placeholder="placeholder"
           :value="modelValue" :min="min" :max="max"
           @input="$emit('update:modelValue', $event.target.value)"
           :autocomplete="name + '-new'" :required="required"
           class="form-control">
    <div class="form-text">{{help}}</div>
  </div>
</template>

<style scoped>

</style>