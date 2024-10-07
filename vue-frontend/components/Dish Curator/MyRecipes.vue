<template>
  <div class="w-full mt-6">
    <div class="flex gap-1 mx-2 flex-wrap justify-between">
      <template v-if="data.length > 0">
        <RecipeCard
          v-for="recipe in data"
          :key="recipe.id"
          :id="recipe.id"
          :title="recipe.title"
          :category="recipe.category"
          :preparation_time="recipe.preparation_time"
          :description="recipe.description"
          :isAdmin="true"
          :url="recipe.featured_images[0].image_url"
          @edit="handleEdit"
          @delete="handleDelete"
        />
      </template>
      <template v-else>
        <div class="w-full text-center text-gray-500 mt-10">
          <p class="text-lg font-semibold">No Listed Recipes</p>
          <p class="mt-2">It looks like you haven't added any recipes yet.</p>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import RecipeCard from "../recipe/RecipeCard.vue";
import { GET_MY_RECIPES } from "@/graphql/queries/creator";
import { DELETE_RECIPES_MUTATION } from "@/graphql/mutations/creator";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { ref, watch } from "vue";

export default {
  name: "MyRecipes",
  components: {
    RecipeCard,
  },
  setup() {
    const data = ref([]);
    const { result, refetch } = useQuery(GET_MY_RECIPES);

    watch(result, (newResult) => {
      if (newResult && newResult.recipes) {
        data.value = newResult.recipes;
      }
    });

    const { mutate: deleteRecipe } = useMutation(DELETE_RECIPES_MUTATION, {
      update() {
        refetch();
      },
    });

    const handleEdit = (recipeId) => {
      console.log("Editing recipe with ID:", recipeId);
    };

    const handleDelete = async (recipeId) => {
      try {
        await deleteRecipe({ recipe_id: parseInt(recipeId) });
      } catch (error) {
        console.error("Error deleting the recipe:", error);
      }
    };

    return {
      data,
      handleEdit,
      handleDelete,
    };
  },
};
</script>

<style scoped>
.text-center {
  text-align: center;
}
</style>
