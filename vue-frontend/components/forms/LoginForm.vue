<script setup>
import { ref } from "vue";
import { Field, Form, ErrorMessage } from "vee-validate";
import * as Yup from "yup";
import MessageDisplay from "../ui/MessageDisplay.vue";
import { useUserStore } from "@/store/userStore"; // Ensure this path is correct
import { LOGIN_USER_MUTATION } from "@/graphql/mutations/creator"; // Adjust the path to your mutation
import { useMutation } from "@vue/apollo-composable";
import { useRouter } from "vue-router"; // Import useRouter for navigation

// Access the user store
const store = useUserStore();
const router = useRouter(); // Initialize the router

const email = ref("");
const password = ref("");
const message = ref("");

// Define the validation schema using Yup
const schema = Yup.object({
  email: Yup.string().email("Invalid email").required("Field is required"),
  password: Yup.string().required("Field is required"),
});

// Define the login function
const login = async () => {
  try {
    const { mutate } = useMutation(LOGIN_USER_MUTATION, {
      context: { isPublicQuery: true },
    });

    const response = await mutate({
      email: email.value,
      password: password.value,
    });

    message.value = "Directing to dashboard";
    store.setUserId(response.data.loginUser.id);
    store.setToken(response.data.loginUser.token);
    router.push({
      name: "curator-id",
      params: { id: response.data.loginUser.id },
    });
  } catch (err) {
    message.value =
      "An error occurred during login. Either email or password is incorrect.";
  }
};
</script>

<template>
  <section class="bg-gray-50 dark:bg-gray-900">
    <div
      class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0"
    >
      <div
        class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700"
      >
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <MessageDisplay :message="message" />
          <h1
            class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white"
          >
            Log In to your account
          </h1>
          <Form
            v-slot="{ errors }"
            @submit="login"
            :validation-schema="schema"
            class="space-y-4 md:space-y-6"
          >
            <label
              for="email"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Your email</label
            >
            <Field
              name="email"
              v-model="email"
              placeholder="email@example.com"
              class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
            <ErrorMessage name="email" class="text-red-500" />

            <label
              for="password"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Password</label
            >
            <Field
              name="password"
              type="password"
              v-model="password"
              placeholder="........"
              class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
            <ErrorMessage name="password" class="text-red-500" />

            <button
              type="submit"
              class="block text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700"
            >
              Log In
            </button>
            <h1
              class="font-bold leading-tight tracking-tight text-gray-900 dark:text-white"
            >
              <nuxt-link to="/signup">Back to signup</nuxt-link>
            </h1>
          </Form>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped></style>
