<template>
  <div class="flex gap-2">
    <button class="text-white font-bold py-2 px-4 rounded-full flex gap-2 mt-1">
      <i
        v-for="(star, index) in 5"
        :key="index"
        class="fa fa-star"
        aria-hidden="true"
        :class="getStarClass(index)"
        @click="changeRating(index + 1)"
      ></i>
    </button>

    <p class="mt-1.5 text-xl">({{ data.count }})</p>
  </div>
</template>

<script>
import { useQuery, useMutation } from "@vue/apollo-composable";
import { ref, watch } from "vue";
import { useRoute } from "nuxt/app";
import { useUserStore } from "~/store/userStore";
import {
  GET_ALL_RECIPE_RATINGS,
  CHECK_USER_RATED_RECIPE,
} from "@/graphql/queries/recipes";
import {
  ADD_RATING_MUTATION,
  REMOVE_RATING_MUTATION,
} from "@/graphql/mutations/recipes";

export default {
  name: "RatingRecipe",
  setup() {
    const route = useRoute();
    const store = useUserStore();
    const recipeId = route.params.id;
    const userId = store.userId; // Replace with actual logged-in user info.
    const data = ref({ count: 0 });
    const currentRating = ref(0);

    // Fetch recipe ratings
    const { result, refetch } = useQuery(GET_ALL_RECIPE_RATINGS, {
      recipe_id: recipeId,
    });

    watch(result, (newResult) => {
      if (newResult && newResult.ratings_aggregate.aggregate) {
        const avgRating = newResult.ratings_aggregate.aggregate.avg.rating;
        currentRating.value = Math.round(avgRating);
        data.value.count = avgRating.toFixed(2);
      }
    });

    // Add and remove rating mutation functions
    const { mutate: addRating } = useMutation(ADD_RATING_MUTATION);
    const { mutate: removeRating } = useMutation(REMOVE_RATING_MUTATION);

    // Track if the user has rated the recipe
    const { result: userRatingResult } = useQuery(CHECK_USER_RATED_RECIPE, {
      recipe_id: recipeId,
      user_id: userId,
    });

    // Watch user rating results
    watch(userRatingResult, (newResult) => {
      if (newResult && newResult.ratings_aggregate) {
        const hasRated = newResult.ratings_aggregate.aggregate
          ? newResult.ratings_aggregate.aggregate.count > 0
          : false;
        currentRating.value = hasRated ? currentRating.value : 0; // Reset rating if not rated
      }
    });

    // Function to update rating when user clicks a star
    const changeRating = async (rating) => {
      try {
        // Check if the user has already rated
        const userRatingData = userRatingResult.value; // Get the latest user rating result

        if (userRatingData && userRatingData.ratings_aggregate) {
          const userAggregate = userRatingData.ratings_aggregate.aggregate;

          if (userAggregate && userAggregate.count > 0) {
            // If already rated, remove the old rating
            await removeRating({
              recipe_id: recipeId,
              user_id: userId,
            });
          }
        }

        // Add the new rating
        await addRating({
          recipe_id: recipeId,
          user_id: userId,
          rating: rating,
        });

        // Refetch ratings to update UI
        await refetch();
      } catch (error) {
        console.error("Error updating rating:", error);
      }
    };

    // Get the class for each star (filled or unfilled)
    const getStarClass = (index) => {
      return index < currentRating.value ? "text-yellow-400" : "text-gray-400";
    };

    return {
      data,
      getStarClass,
      changeRating,
      userId,
    };
  },
};
</script>

<style scoped>
.fa-star {
  cursor: pointer;
  font-size: 1.5rem;
}
</style>
