<script>
export default {
  name: "SelectInput",
  props: {
    items: Array,
    required: String,
    name: String,
    label: String,
  }
}
</script>

<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{label}}</label>
    <select :id="name" class="form-select" :name="name" :required="required">
      <option v-for="option in items" :value="option.value" :key="option.value">
        {{option.text}}
      </option>
    </select>
  </div>
</template>

<style scoped>

</style>