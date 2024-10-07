<template>
  <div class="max-w-full md:max-w-2xl mx-auto mt-7 relative px-4">
    <div class="mb-5 text-left">
      <h1
        class="font-mono text-2xl sm:text-xl text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
      >
        {{ title }}
      </h1>
      <h1
        class="font-mono text-base sm:text-lg text-gray-700 rounded md:border-0 md:p-0 dark:text-gray-400"
      >
        {{ description }}
      </h1>
      <div class="mt-3 flex gap-4 flex-col sm:flex-row">
        <LikeRecipe />
        <RatingRecipe />
      </div>
      <Ingrident />
    </div>

    <div id="controls-carousel" class="relative w-full" data-carousel="static">
      <!-- Carousel wrapper -->
      <div class="relative h-56 overflow-hidden rounded-lg md:h-96 mt-4">
        <!-- Show a spinner when loading -->
        <Spinner v-if="loading" />

        <!-- Display an error message if there's an error -->
        <div v-if="error" class="text-red-500">
          {{ error.message }}
        </div>

        <!-- Loop through the fetched data and display images -->
        <div
          v-for="(image, index) in data"
          :key="index"
          :class="[
            'absolute block rounded-lg w-full transition duration-700 ease-in-out',
            { hidden: index !== activeImage, flex: index === activeImage },
          ]"
        >
          <img
            :src="image"
            class="object-cover rounded-lg w-full h-full"
            alt="Image"
            style="object-fit: cover"
          />
        </div>
      </div>

      <!-- Slider controls -->
      <button
        type="button"
        class="absolute top-0 left-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
        @click="prevImage"
      >
        <span
          class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 group-hover:bg-white/50 group-focus:ring-4 group-focus:ring-white"
        >
          <svg
            class="w-4 h-4 text-white"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 6 10"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M5 1L1 5l4 4"
            />
          </svg>
          <span class="sr-only">Previous</span>
        </span>
      </button>

      <button
        type="button"
        class="absolute top-0 right-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
        @click="nextImage"
      >
        <span
          class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 group-hover:bg-white/50 group-focus:ring-4 group-focus:ring-white"
        >
          <svg
            class="w-4 h-4 text-white"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 6 10"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M1 9l4-4-4-4"
            />
          </svg>
          <span class="sr-only">Next</span>
        </span>
      </button>
    </div>
    <RecipeSteps />
    <CommentRecipe />
  </div>
</template>

<script>
import { GET_ALL_RECIPES_IMAGES } from "@/graphql/queries/recipes";
import { useQuery } from "@vue/apollo-composable";
import { ref, watch } from "vue";
import Spinner from "../ui/Spinner.vue";
import { useRoute } from "nuxt/app";
import Ingrident from "./Ingrident.vue";
import LikeRecipe from "./LikeRecipe.vue";
import CommentRecipe from "./CommentRecipe.vue";
import RatingRecipe from "./RatingRecipe.vue";
import RecipeSteps from "./RecipeSteps.vue";

export default {
  name: "ImageCard",
  components: {
    Spinner,
    Ingrident,
    LikeRecipe,
    CommentRecipe,
    RatingRecipe,
    RecipeSteps,
  },
  data() {
    return {
      activeImage: 0,
    };
  },
  methods: {
    nextImage() {
      this.activeImage = (this.activeImage + 1) % this.data.length;
    },
    prevImage() {
      this.activeImage =
        (this.activeImage - 1 + this.data.length) % this.data.length;
    },
  },
  setup() {
    const route = useRoute();
    const recipeId = route.params.id;
    const data = ref([]);
    const loading = ref(true);
    const error = ref(null);
    const title = route.query.title; // Accessing the query parameter
    const category = route.query.category; // Accessing the query parameter
    const preparationTime = route.query.preparation_time; // Accessing the query parameter
    const description = route.query.description; // Accessing the query parameter

    // Fetch the images from the GraphQL query
    const {
      result,
      loading: queryLoading,
      error: queryError,
    } = useQuery(
      GET_ALL_RECIPES_IMAGES,
      {
        recipe_id: recipeId,
      },
      {
        context: { isPublicQuery: true },
      }
    );

    // Watch for query results and errors
    watch(result, (newResult) => {
      if (newResult && newResult.featured_images) {
        data.value = newResult.featured_images.map((img) => img.image_url); // Assuming image URLs are in 'url'
      }
    });

    watch(queryLoading, (newLoading) => {
      loading.value = newLoading;
    });

    watch(queryError, (newError) => {
      if (newError) {
        error.value = newError;
      }
    });

    return {
      data,
      loading,
      error,
      title, // Return title
      category, // Return category
      preparationTime, // Return preparationTime
      description, // Return description
    };
  },
};
</script>

<style scoped>
img {
  height: 300px;
  width: 100%;
  object-fit: cover;
}

button {
  outline: none;
}
</style>
