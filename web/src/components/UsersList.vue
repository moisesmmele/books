<script>
import security from "@/components/security";
import notie from "notie/dist/notie";

export default {
  name: "UsersList",

  data() {
    return {
      users: []
    }
  },

  beforeMount() {
    security.requireToken()
    fetch(process.env.VUE_APP_API_URL + "/admin/users", security.requestOptions(""))
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            notie.alert({
              type: "error",
              message: response.message,
            })
          } else {
            this.users =  response.data.users
          }
        })
        .catch((error) => {
          notie.alert({
            type: "error",
            message: error,
          })
        })

  }

}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Users</h1>
        <hr>
        <table class="table table-compact table-striped">
          <thead>
          <tr>
            <th>User</th>
            <th>Email</th>
          </tr>
          </thead>
          <tbody>
            <tr v-for="user in this.users" :key="user.id">
              <td><router-link :to="`/admin/users/${user.id}`">{{user.last_name}}, {{user.first_name}}</router-link></td>
              <td>{{user.email}}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>