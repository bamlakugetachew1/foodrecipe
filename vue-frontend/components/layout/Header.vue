<script>
import { useUserStore } from "@/store/userStore";
import { useRouter } from "vue-router";

export default {
  name: "Header",
  props: {
    isAdmin: Boolean,
    isHomepage: Boolean,
  },
  setup() {
    const store = useUserStore();
    const router = useRouter();
    const isMenuOpen = ref(false);

    const toggleMenu = () => {
      isMenuOpen.value = !isMenuOpen.value;
    };

    const logout = () => {
      store.logout();
      router.push({ name: "login" });
    };

    return {
      isMenuOpen,
      toggleMenu,
      logout,
      id: store.userId,
    };
  },
};
</script>

<template>
  <nav
    class="bg-white border-b border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900 shadow"
  >
    <div class="container flex flex-wrap items-center justify-between mx-auto">
      <a href="#" class="flex items-center">
        <img
          src="https://flowbite.com/docs/images/logo.svg"
          class="h-6 mr-3 sm:h-9"
          alt="FoodIE Logo"
        />
        <span
          class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
          ><nuxt-link to="/">GEBETA </nuxt-link>
        </span>
      </a>
      <button
        @click="toggleMenu"
        type="button"
        class="inline-flex items-center p-2 ml-3 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
        aria-controls="navbar-default"
        aria-expanded="isMenuOpen"
      >
        <span class="sr-only">Open main menu</span>
        <svg
          class="w-6 h-6"
          aria-hidden="true"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fill-rule="evenodd"
            d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
            clip-rule="evenodd"
          ></path>
        </svg>
      </button>
      <div
        :class="{ hidden: !isMenuOpen, block: isMenuOpen }"
        class="w-full md:block md:w-auto"
        id="navbar-default"
      >
        <ul
          class="flex flex-col p-4 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700"
        >
          <li v-if="isHomepage && id == null">
            <a
              class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
              href="#"
            >
              <nuxt-link to="/signup">SignUp</nuxt-link>
            </a>
          </li>
          <li v-if="isHomepage && id == null">
            <a
              class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
              href="#"
            >
              <nuxt-link to="/login">Login</nuxt-link>
            </a>
          </li>
          <li
            v-if="!isHomepage"
            class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
          >
            <nuxt-link :to="`/curator/${id}`">My Recipes</nuxt-link>
          </li>
          <li v-if="!isHomepage">
            <a
              class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
              href="#"
            >
              <nuxt-link to="/curator/add">Add Recipes</nuxt-link>
            </a>
          </li>
          <li
            v-if="!isHomepage || id !== null"
            class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
          >
            <button @click="logout">Logout</button>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scoped></style>
