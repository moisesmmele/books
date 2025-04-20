<template>
  <AppHeader />
  <div>
    <RouterView @success="success" @error="error" @warning="warning"/>
  </div>
  <AppFooter />
</template>
<script>
import AppHeader from "@/components/AppHeader.vue";
import {store} from "@/components/store.js";
import AppFooter from "@/components/AppFooter.vue";
import notie from "notie/dist/notie";

// const getCookie = (name) => {
//   return document.cookie.split(";")
//       .reduce((r, v) => {
//         const parts = v.split("=");
//         return parts[0] === name ? decodeURIComponent(parts[1]) : r;
//       }, "")
// }

const getCookie = (name) => {
  const cookies = document.cookie.split("; ");
  for (let cookie of cookies) {
    const [key, ...valParts] = cookie.split("=");
    if (key === name) {
      return decodeURIComponent(valParts.join("="));
    }
  }
  return null;
};



export default {
  name: 'App',
  components: {
    AppHeader,
    AppFooter,
  },
  data() {
    return {
      store,
    }
  },
  methods: {
    success(msg) {
      notie.alert({
        type: "success",
        text: msg
      })
    },
    error(msg) {
      notie.alert({
        type: "error",
        text: msg
      })
    },
    warning(msg) {
      notie.alert({
        type: "warning",
        text: msg
      })
    }
  },
  beforeMount() {
  let data = getCookie("_site_data")
    if (data !== null) {
      let cookieData = JSON.parse(data)
      store.token = cookieData.token.token
      store.user = {
        id: cookieData.user.id,
        first_name: cookieData.user.first_name,
        last_name: cookieData.user.last_name,
        email: cookieData.user.email,
      }
    }
  },
  mounted() {
    console.log("App mounted.");
  }
}
</script>
<style>
</style>