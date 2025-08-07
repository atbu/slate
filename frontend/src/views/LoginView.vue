<script setup lang="ts">
import LoginForm from '@/components/LoginForm.vue';
</script>

<script lang="ts">
export default {
    name: 'login'
};
</script>

<template>
    <h1>Login page</h1>
    <RouterLink to="/">Go back</RouterLink>
    <LoginForm />
</template>