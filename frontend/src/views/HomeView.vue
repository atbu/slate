<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

import Button from "primevue/button";

onMounted(async () => {
    await loadData();
});

const loadData = async () => {
    await axios.get("http://localhost:8080/api/auth/currentuser", { withCredentials: true }).then((response) => {
            currentUserUsername.value = response.data.username;
        }).catch((error) => {
            if(error.response.status === 401) {
                currentUserUsername.value = "";
                return;
            }

            throw error;
    });
}

const logout = async () => {
    await axios.post("http://localhost:8080/api/auth/logout", null, { withCredentials: true }).then(() => {
        loadData();
    });
}

const currentUserUsername = ref("");
</script>

<script lang="ts">
export default {
    name: 'home'
};
</script>

<template>
    <h1>Home page</h1>
    <RouterLink to="/login">Login</RouterLink>
    <br />
    <RouterLink to="/register">Register</RouterLink>
    <h1>{{ currentUserUsername != "" ? `Logged in as ${currentUserUsername}` : "Not logged in" }}</h1>
    <Button v-if="currentUserUsername != ''" label="Logout" @click="logout" />
</template>