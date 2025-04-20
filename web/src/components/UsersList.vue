<script>
import security from "@/components/security";
import notie from "notie/dist/notie";
import {store} from "@/components/store";

export default {
  name: "UsersList",

  data() {
    return {
      users: [],
      ready: false,
      store: store
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
            this.ready = true
          }
        })
        .catch((error) => {
          notie.alert({
            type: "error",
            message: error,
          })
        })

  },
  logoutUser(id) {
    if (id !== store.user.id) {
      notie.confirm({
        text: "Are you sure you want to logout the user?",
        submitText: "Log out",
        submitCallback: function () {
          console.log(`logging user ${id} out`)
        }
      })
    } else {
      this.$emit("error", "Can't log yourself out")
    }
  }
}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">Users</h1>
        <hr>
        <table v-if="this.ready" class="table table-compact table-striped">
          <thead>
          <tr>
            <th>User</th>
            <th>Email</th>
            <th>Status</th>

          </tr>
          </thead>
          <tbody>
            <tr v-for="user in this.users" :key="user.id">
              <td><router-link :to="`/admin/users/${user.id}`">{{user.last_name}}, {{user.first_name}}</router-link></td>
              <td>{{user.email}}</td>
              <td v-if="user.token.id > 0">
                <span class="badge bg-success" @click="logoutUser(user.id)">Logged in</span>
              </td>
              <td v-else>
                <span class="badge bg-danger">Not logged in</span>
              </td>

            </tr>
          </tbody>
        </table>
        <div v-else class="container">Loading...</div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>