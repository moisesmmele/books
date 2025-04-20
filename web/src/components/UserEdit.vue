<script>
import security from "@/components/security";
import {store} from "@/components/store";
import FormTag from "@/components/forms/FormTag.vue";
import textInput from "@/components/forms/TextInput.vue";
import notie from "notie/dist/notie";

export default {
  name: "UserEdit",
  computed: {
    store() {
      return store
    }
  },

  beforeMount() {
    security.requireToken()
    if (parseInt(String(this.$route.params.userId), 10) > 0) {
      //editing existing user
    }
  },

  mounted() {
    console.log("this.user.id: ", parseInt(String(this.user.id), 10))
    console.log("store.user.id: ", store.user.id)
  },

  data() {
    return {
      user: {
        id: this.$route.params.userId,
        first_name: "",
        last_name: "",
        email: "",
        password: "",
      }
    }
  },

  components: {
    textInput,
    FormTag,
  },

  methods: {
    submitHandler() {
      const payload = {
        id: parseInt(String(this.$route.params.userId), 10),
        first_name: this.user.first_name,
        last_name: this.user.last_name,
        email: this.user.email,
        password: this.user.password
      }
      fetch(process.env.VUE_APP_API_URL + "/admin/users/save", security.requestOptions(payload))
          .then((response) => response.json())
          .then((response) => {
            if (response.error) {
              notie.alert({
                type: "error",
                text: response.message
              })
            } else {
              notie.alert({
                type: "success",
                text: response.message
              })
            }
          })
          .catch((error) => {
            notie.alert({
              type: "error",
              text: error
            })
          })
    },
    confirmDelete() {

    }
  }
}

</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Edit User</h1>
        <hr>
        <form-tag @userEditEvent="submitHandler" name="user-form" event="userEditEvent">
          <text-input v-model="user.first_name" type="text" required="true" label="First Name"
                      value="first-name" name="first-name"/>
          <text-input v-model="user.last_name" type="text" required="true" label="Last Name"
                      value="last-name" name="last-name"/>
          <text-input v-model="user.email" type="email" required="true" label="Email"
                      value="email" name="email"/>
          <text-input v-if="this.user.id === 0" v-model="user.password" type="password" required="true"
                      label="password" value="password" name="password"/>
          <text-input v-else v-model="user.password" type="password" required="false"
                      label="password" value="password" name="password"/>
          <div class="float-start">
            <input type="submit" class="btn btn-primary me-2" value="Save"/>
            <router-link to="/admin/users" class="btn btn-secondary">Cancel</router-link>
          </div>
          <div class="float-end">
            <a v-if="this.$route.params.userId > 0 && (parseInt(String(this.user.id), 10) !== store.user.id)"
            class="btn btn-danger" href="javascript:void(0);" @click="confirmDelete(this.user.id)">Delete</a>
          </div>
          <div class="clearfix"></div>
        </form-tag>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>