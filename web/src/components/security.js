import {store} from "@/components/store";
import router from "@/router/index";

let security = {
    requireToken: function() {
        if (store.token === null) {
            router.push("/login");
            return false;
        }
    },
    requestOptions: function(payload) {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + store.token);

        return {
            method: "POST",
            headers: headers,
            body: JSON.stringify(payload)
        }
    }
}

export default security;