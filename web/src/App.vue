<template>
  <AppHeader />
  <div>
    <RouterView/>
  </div>
  <AppFooter />
</template>
<script>
import AppHeader from "@/components/AppHeader.vue";
import {store} from "@/components/store.js";
import AppFooter from "@/components/AppFooter.vue";

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
  beforeMount() {
  let data = getCookie("_site_data")
    console.log("data: " + typeof data + data)
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
}
</script>
<style>
</style>