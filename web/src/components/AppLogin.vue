<script>
import TextInput from "@/components/forms/TextInput.vue";
import FormTag from "@/components/forms/FormTag.vue"
import {store} from "@/components/store.js";
import router from "@/router";
import notie from "notie/dist/notie";

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
          .then((response) => response.json())
          .then((data) => {
            if (data.error) {
              console.log(data.message)
              notie.alert({
                type: "error",
                text: data.message,
              })
            } else {
              store.token = data.data.token.token
              store.user = {
                id: data.data.user.id,
                first_name: data.data.user.first_name,
                last_name: data.data.user.last_name,
                email: data.data.user.email,
              }
              let date = new Date();
              let expDays = 1;
              date.setTime(date.getTime() + (expDays * 24 * 60 * 60 * 1000));
              document.cookie = "_site_data="
                  + JSON.stringify(data.data) + "; "
                  + "expires=" + date.toUTCString() + "; "
                  + "path=/; "
                  + "SameSite=Strict; "
                  + "Secure; " ;
              router.push("/")
            }
          })
    }
  },
  data() {
    return {
      email: "",
      password: "",
      store,
    }
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