<template>
  <div>
    <h1 class="font-serif text-lg text-gray-600 mt-4 ml-2 mb-2">
      Search Results for "{{ searchQuery }}"
    </h1>

    <!-- Loading state -->
    <div v-if="loading">
      <Spinner />
    </div>

    <!-- Error state -->
    <div v-if="error" class="text-center text-red-500">
      <h1 class="font-serif text-lg text-gray-600 mt-4 ml-2 mb-2">
        Error loading recipes: {{ error.message }}
      </h1>
    </div>

    <!-- Recipes display -->
    <div v-if="data && data.length">
      <div class="w-full">
        <div class="flex gap-1 mx-2 flex-wrap justify-between">
          <RecipeCard
            v-for="recipe in data"
            :key="recipe.id"
            :id="recipe.id"
            :title="recipe.title"
            :category="recipe.category"
            :preparation_time="recipe.preparation_time"
            :description="recipe.description"
            :isAdmin="false"
            :url="recipe.featured_images[0].image_url"
          />
        </div>
      </div>
    </div>

    <!-- No results found -->
    <div v-else-if="!loading && !data.length">
      <NoRecipesFound />
    </div>
  </div>
</template>

<script>
import { BROWSEBY_TITLE_CATEGORY_OR_CATAGORY } from "@/graphql/queries/recipes";
import { useQuery } from "@vue/apollo-composable";
import { ref, watch } from "vue";
import RecipeCard from "./RecipeCard.vue";
import Spinner from "../ui/Spinner.vue";
import NoRecipesFound from "./NoRecipesFound.vue";

export default {
  name: "SearchRecipes",
  components: {
    RecipeCard,
    Spinner,
    NoRecipesFound,
  },
  props: {
    searchQuery: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const data = ref([]);
    const loading = ref(true); // Start with loading state as true
    const error = ref(null);

    // Fetching the query result
    const {
      result,
      loading: queryLoading,
      error: queryError,
    } = useQuery(
      BROWSEBY_TITLE_CATEGORY_OR_CATAGORY,
      {
        search: props.searchQuery,
      },
      {
        context: { isPublicQuery: true },
      }
    );

    // Watch the query loading state
    watch(queryLoading, (newLoading) => {
      loading.value = newLoading; // Set loading to query loading state
    });

    // Watch the result and handle data or error states
    watch(result, (newResult) => {
      console.log(result);
      if (newResult && newResult.recipes) {
        data.value = newResult.recipes;
      }
    });

    // Watch for query errors
    watch(queryError, (newError) => {
      if (newError) {
        error.value = newError;
      }
    });

    return {
      data,
      loading,
      error,
    };
  },
};
</script>

<style scoped>
/* Add your custom styles if needed */
</style>
