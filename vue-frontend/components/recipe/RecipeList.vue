<template>
  <RecipeFilter @filter-change="handleFilterChange" />
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
    <a
      href="#"
      @click.prevent="loadMore"
      class="mb-3 ml-1 content-center inline-flex items-center px-3 py-2 mt-3 text-sm font-medium text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
    >
      Load More
    </a>
  </div>
</template>

<script>
import { ref, watchEffect } from "vue";
import RecipeCard from "./RecipeCard.vue";
import { useQuery } from "@vue/apollo-composable";
import {
  RETRIVE_ALL_RECIPES,
  FILTER_BY_TIME,
  FILTER_BY_TITLE,
} from "@/graphql/queries/recipes";
import RecipeFilter from "./RecipeFilter.vue";

export default {
  name: "RecipeList",
  components: {
    RecipeCard,
    RecipeFilter,
  },
  setup() {
    const pages = ref(1);
    const data = ref([]);
    const currentQuery = ref(RETRIVE_ALL_RECIPES);

    // Function to fetch recipes based on the current query and page
    const fetchRecipes = () => {
      const { result, loading, error } = useQuery(
        currentQuery.value,
        {
          limit: 6 * pages.value,
        },
        {
          context: { isPublicQuery: true },
        }
      );

      watchEffect(() => {
        if (result.value && result.value.recipes) {
          data.value = result.value.recipes;
        }
      });

      return { loading, error };
    };

    // Initial fetch of recipes
    const { loading, error } = fetchRecipes();

    // Function to load more recipes
    const loadMore = () => {
      pages.value++;
      fetchRecipes(); // Fetch recipes again with the updated page
    };

    const handleFilterChange = (filter) => {
      const queryMap = {
        all: RETRIVE_ALL_RECIPES,
        time: FILTER_BY_TIME,
        title: FILTER_BY_TITLE,
      };

      currentQuery.value = queryMap[filter] || RETRIVE_ALL_RECIPES;
      pages.value = 1;
      fetchRecipes(); // Fetch recipes again with the updated query
    };

    return {
      data,
      loadMore,
      handleFilterChange,
      loading,
      error,
    };
  },
};
</script>

<style scoped></style>
