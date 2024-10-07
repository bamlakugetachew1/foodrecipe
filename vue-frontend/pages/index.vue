<template>
  <Header :isHomepage="true" />
  <RecipesIndex />
</template>

<script>
import Header from "~/components/layout/Header.vue";
import RecipesIndex from "./recipes/index.vue";

export default {
  name: "HomePage",
  components: {
    Header,
    RecipesIndex,
  },
};
</script>
