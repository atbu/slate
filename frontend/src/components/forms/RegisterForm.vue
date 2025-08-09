<template>
    <Form v-slot="$form" :initialValues @submit="onFormSubmit" class="flex flex-col gap-4 w-full sm:w-60">
        <div class="flex flex-col gap-1">
            <InputText v-model="email" name="email" type="text" placeholder="Email" fluid />
            <Message v-if="$form.email?.invalid" severity="error" size="small" variant="simple">{{ $form.email.error.message }}</Message>
        </div>
        <div class="flex flex-col gap-1">
            <InputText v-model="username" name="username" type="text" placeholder="Username" fluid />
            <Message v-if="$form.username?.invalid" severity="error" size="small" variant="simple">{{ $form.username.error.message }}</Message>
        </div>
        <div class="flex flex-col gap-1">
            <Password v-model="password" name="password" placeholder="Password" :feedback="false" toggleMask fluid />
            <Message v-if="$form.password?.invalid" severity="error" size="small" variant="simple">
                <ul class="my-0 px-4 flex flex-col gap-1">
                    <li v-for="(error, index) of $form.password.errors" :key="index">{{ error.message }}</li>
                </ul>
            </Message>
        </div>
        <Button type="submit" severity="secondary" label="Submit" />
    </Form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import { Form } from "@primevue/forms";
import InputText from "primevue/inputtext";
import Message from "primevue/message";
import Password from "primevue/password";
import Button from "primevue/button";

const router = useRouter();

const initialValues = ref([]);

const email = ref("");
const username = ref("");
const password = ref("");

const onFormSubmit = async () => {
    await axios.post("http://localhost:8080/api/auth/register", {
        email: email.value,
        username: username.value,
        password: password.value
    }, { withCredentials: true }).then(() => {
        router.push("/");
    });
};
</script>