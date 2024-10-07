<template>
  <div class="flex gap-2">
    <button
      @click="toggleLike"
      class="bg-red-500 text-white font-bold py-2 px-4 rounded-full"
      :disabled="!userId"
    >
      <i class="fa fa-heart" aria-hidden="true" id="space"></i>
    </button>
    <p class="mt-1.5 text-xl">({{ data.count }})</p>
  </div>
</template>

<script>
import {
  GET_ALL_RECIPES_LIKES,
  CHECK_USER_LIKED_RECIPE,
} from "@/graphql/queries/recipes";
import {
  LIKE_RECIPE_MUTATION,
  UNLIKE_RECIPE_MUTATION,
} from "@/graphql/mutations/recipes";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { ref, watch } from "vue";
import { useRoute } from "nuxt/app";
import { useUserStore } from "~/store/userStore";

export default {
  name: "LikeRecipe",
  setup() {
    const route = useRoute();
    const store = useUserStore();
    const recipeId = route.params.id;

    const data = ref({ count: 0 });
    const userId = store.userId;

    const { result, refetch } = useQuery(
      GET_ALL_RECIPES_LIKES,
      {
        recipe_id: recipeId,
      },
      {
        context: { isPublicQuery: true },
      }
    );

    watch(result, (newResult) => {
      if (newResult && newResult.likes_aggregate.aggregate) {
        data.value.count = newResult.likes_aggregate.aggregate.count;
      }
    });

    const { mutate: likeRecipeMutate } = useMutation(LIKE_RECIPE_MUTATION);
    const { mutate: unlikeRecipeMutate } = useMutation(UNLIKE_RECIPE_MUTATION);

    const toggleLike = async () => {
      if (!userId) {
        console.log("User not logged in. Please log in to like the recipe.");
        return;
      }

      try {
        const { result: likedResult } = useQuery(CHECK_USER_LIKED_RECIPE, {
          recipe_id: recipeId,
          user_id: userId,
        });

        const liked =
          (await likedResult.value?.likes_aggregate?.aggregate?.count) > 0;

        if (liked) {
          await unlikeRecipeMutate({
            recipe_id: recipeId,
            user_id: userId,
          });
        } else {
          await likeRecipeMutate({
            recipe_id: recipeId,
            user_id: userId,
          });
        }

        refetch();
      } catch (error) {
        console.error("Error toggling like:", error);
      }
    };

    return {
      data,
      toggleLike,
      userId,
    };
  },
};
</script>
