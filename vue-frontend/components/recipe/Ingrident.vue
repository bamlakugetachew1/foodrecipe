<template>
  <div class="flex gap-3 mt-2">
    <h1
      v-for="(Ingrident, index) in data"
      :key="index"
      class="font-mono font-semibold text-gray-700"
    >
      {{ Ingrident.ingredient_name }}
      {{ Ingrident.quantity }}
    </h1>
  </div>
</template>

<script>
import { GET_ALL_RECIPES_INGRIDENTS } from "@/graphql/queries/recipes";
import { useQuery } from "@vue/apollo-composable";
import { ref, watch } from "vue";
import { useRoute } from "nuxt/app";

export default {
  name: "Ingrident",
  setup() {
    const route = useRoute();
    const recipeId = route.params.id;
    const data = ref([]);
    const { result } = useQuery(
      GET_ALL_RECIPES_INGRIDENTS,
      {
        recipe_id: recipeId,
      },
      {
        context: { isPublicQuery: true },
      }
    );

    watch(result, (newResult) => {
      if (newResult && newResult.Ingredients) {
        data.value = newResult.Ingredients;
      }
    });

    return {
      data,
    };
  },
};
</script>

<style scoped></style>
