<template>
  <Header />
  <div>
    <MyRecipes />
  </div>
</template>

<script>
import MyRecipes from "~/components/Dish Curator/MyRecipes.vue";
import Header from "~/components/layout/Header.vue";

export default {
  components: {
    MyRecipes,
    Header,
  },
};
</script>
