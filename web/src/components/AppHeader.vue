<script>
import {store} from "@/components/store.js";
import router from "@/router";
import security from "@/components/security";

export default {
  name: "AppHeader",
  data() {
    return {
      store,
    }
  },
  methods: {
    logout() {
      const payload = {
        token: store.token,
      }

      fetch(process.env.VUE_APP_API_URL + "/users/logout", security.requestOptions(payload))
          .then(res => res.json())
          .then(data => {
            if (data.error) {
              console.log(data.message);
            } else {
              store.token = null
              store.user = {}
              document.cookie = "_site_data=null; path=/; SameSite=Strict; Secure; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
              router.push({ path: "/login" })
            }
          })
    },
  }
}
</script>

<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Navbar</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link class="nav-link" aria-current="page" to="/">Home</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/books">Books</router-link>
          </li>
          <li v-if="store.token !== null" class="nav-item dropdown">
            <a href="#" class="nav-link dropdown-toggle" id="navbar-dropdown" role="button"
               data-bs-toggle="dropdown" aria-expanded="false">
              Admin
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbar-dropdown">
              <li>
                <router-link class="dropdown-item" to="/admin/users">Manage Users</router-link>
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/users/0">Add User</router-link>
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/books">Manage Books</router-link>
              </li>
              <li>
                <router-link class="dropdown-item" to="/admin/books/0">Add Book</router-link>
              </li>
            </ul>
          </li>
          <li class="nav-item">
            <router-link v-if="store.token == null" class="nav-link" aria-current="page" to="/login">Login</router-link>
            <router-link v-else class="nav-link" aria-current="page" to="javascript:void(0);" @click="logout">Logout</router-link>
          </li>
        </ul>
        <span class="navbar-text">
          {{store.user.first_name ?? ''}}
        </span>
      </div>
    </div>
  </nav>
</template>

<style scoped>

</style>