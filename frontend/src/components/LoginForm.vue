<template>
    <Form v-slot="$form" :initialValues @submit="onFormSubmit" class="flex flex-col gap-4 w-full sm:w-60">
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
import { Form } from "@primevue/forms";
import InputText from "primevue/inputtext";
import Message from "primevue/message";
import Password from "primevue/password";
import Button from "primevue/button";
import { ref } from "vue";

const onFormSubmit = () => {
    console.log(`Submitted form, username ${username.value} password ${password.value}`);
}

const initialValues = ref([]);

const username = ref("");
const password = ref("");
</script>
