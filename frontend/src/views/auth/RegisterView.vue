<template>
    <h1>Register page</h1>
    <RouterLink to="/">Go back</RouterLink>
    <RegisterForm />
</template>

<script setup lang="ts">
import RegisterForm from "@/components/forms/RegisterForm.vue";
</script>

<script lang="ts">
export default {
    name: "register"
};
</script>