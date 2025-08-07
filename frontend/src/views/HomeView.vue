<script setup lang="ts">
import { onMounted, ref } from "vue";
import axios from "axios";

onMounted(async () => {
await axios.get("http://localhost:8080/api/auth/currentuser", { withCredentials: true }).then((response) => {
    currentUserUsername.value = response.data.username;
}).catch((error) => {
    if(error.response.status === 401) {
        return;
    }

    throw error;
});
});

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
    <h1>{{ currentUserUsername != "" ? `Logged in as ${currentUserUsername}` : "Not logged in" }}</h1>
</template>