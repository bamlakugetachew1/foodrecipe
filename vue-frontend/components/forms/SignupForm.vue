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
            Sign Up for a new account
          </h1>
          <Form
            v-slot="{ errors }"
            @submit="signup"
            :validation-schema="schema"
            class="space-y-4 md:space-y-6"
          >
            <label
              for="name"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Your name</label
            >
            <Field
              name="name"
              v-model="name"
              placeholder="bamlaku"
              class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            />
            <ErrorMessage name="name" class="text-red-500" />

            <label
              for="email"
              class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
              >Your email</label
            >
            <Field
              name="email"
              v-model="email"
              placeholder="email"
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
              Sign Up
            </button>
            <h1
              class="font-bold leading-tight tracking-tight text-gray-900 dark:text-white"
            >
              <nuxt-link to="/login">Back to login</nuxt-link>
            </h1>
          </Form>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { ref } from "vue";
import { Field, Form, ErrorMessage } from "vee-validate";
import * as Yup from "yup";
import MessageDisplay from "../ui/MessageDisplay.vue";
import { REGISTER_USER_MUTATION } from "@/graphql/mutations/creator"; // Adjust the path if necessary
import { useMutation } from "@vue/apollo-composable";

export default {
  name: "signup",
  components: {
    Field,
    Form,
    ErrorMessage,
    MessageDisplay,
  },
  setup() {
    const name = ref("");
    const email = ref("");
    const password = ref("");
    const message = ref("");

    const { mutate } = useMutation(REGISTER_USER_MUTATION, {
      context: { isPublicQuery: true },
    });

    const schema = Yup.object({
      name: Yup.string().required("Field is required"),
      email: Yup.string().email("Invalid email").required("Field is required"),
      password: Yup.string().required("Field is required"),
    });

    const signup = async () => {
      try {
        const response = await mutate({
          name: name.value,
          email: email.value,
          password: password.value,
        });
        message.value = response.data.register.message;
      } catch (err) {
        message.value = "An error occurred during signup.";
      }
    };

    return {
      name,
      email,
      password,
      message,
      schema,
      signup, // Return the signup method
    };
  },
};
</script>

<style scoped>
/* Add any scoped styles here */
</style>
