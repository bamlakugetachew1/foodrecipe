<template>
  <div>
    <SearchRecipes :searchQuery="searchQuery" />
  </div>
</template>

<script>
import SearchRecipes from "~/components/recipe/SearchRecipes.vue";
import { useRoute } from "vue-router";

export default {
  components: {
    SearchRecipes,
  },
  setup() {
    const route = useRoute();
    const searchQuery = route.query.searchQuery || "";

    return {
      searchQuery,
    };
  },
};
</script>

<style scoped></style>
