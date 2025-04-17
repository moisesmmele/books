<script>
import TextInput from "@/components/forms/TextInput.vue";
import FormTag from "@/components/forms/FormTag.vue"
export default {
  name: "AppLogin",
  components: {TextInput, FormTag},
  methods: {
    submitHandler() {
      console.log("submitHandler called");
      const payload = {
        email: this.email,
        password: this.password,
      }
      const request = {
        method: "POST",
        body: JSON.stringify(payload),
        headers: {
          "Content-Type": "application/json",
        }
      }
      fetch("http://localhost:8082/users/login", request)
          .then((response) => response.json()
          .then((data) => {
            if (data.error) {
              console.log(data.error)
            } else {
              console.log(data)
            }
          }))
    }
  },
  data() {
    return { email: "", password: "" }
  }
}
</script>
<template>
<div class="container">
  <div class="row">
    <div class="col">
      <h1 class="mt-5">Login</h1>
      <hr>
      <FormTag @submitForm="submitHandler" name="login-form" event="submitForm">
        <TextInput v-model="email" label="email" type="email" placeholder="Email" required=""/>
        <TextInput v-model="password" label="password" type="password" placeholder="Password" required=""/>
        <input type="submit" class="btn btn-primary" value="Login" />
      </FormTag>
    </div>
  </div>
</div>
</template>
<style scoped>
</style>