<template>
  <section class="bg-white dark:bg-gray-900 py-8 lg:py-16 antialiased">
    <div class="max-w-2xl mx-auto px-4">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-lg lg:text-2xl font-bold text-gray-900 dark:text-white">
          Comments ({{ data.count }})
        </h2>
      </div>
      <form class="mb-2" @submit.prevent="postComment">
        <div
          class="py-2 px-4 mb-2 bg-white rounded-lg rounded-t-lg border border-gray-200 dark:bg-gray-800 dark:border-gray-700"
        >
          <textarea
            id="comment"
            v-model="newComment"
            rows="6"
            class="px-0 w-full text-sm text-gray-900 border-0 focus:ring-0 focus:outline-none dark:text-white dark:placeholder-gray-400 dark:bg-gray-800"
            placeholder="Write a comment..."
            required
          ></textarea>
        </div>
        <!-- Button is shown only when the input is not empty -->
        <button
          :disabled="!userId"
          v-show="newComment.trim() !== ''"
          type="submit"
          class="block text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700"
        >
          Post
        </button>
      </form>
      <article
        class="p-6 text-base bg-white border-t border-gray-200 dark:border-gray-700 dark:bg-gray-900"
        v-for="comment in comments"
        :key="comment.id"
      >
        <footer class="flex justify-between items-center mb-2">
          <div class="flex items-center">
            <p class="text-sm text-gray-600 dark:text-gray-400">
              <time pubdate datetime="2022-06-23" title="June 23rd, 2022">{{
                formatDate(comment.created_at)
              }}</time>
            </p>
          </div>
        </footer>
        <p class="text-gray-500 dark:text-gray-400">
          {{ comment.comment_text }}
        </p>
      </article>
    </div>
  </section>
</template>

<script>
import { GET_ALL_RECIPE_COMMENTS } from "@/graphql/queries/recipes";
import { ADD_COMMENTS_MUTATION } from "@/graphql/mutations/recipes";
import { useQuery, useMutation } from "@vue/apollo-composable";
import { ref, watch } from "vue";
import { useRoute } from "nuxt/app";
import { useUserStore } from "~/store/userStore";

export default {
  name: "CommentRecipe",
  setup() {
    const route = useRoute();
    const store = useUserStore();
    const recipeId = route.params.id;
    const data = ref({ count: 0 });
    const comments = ref([]);
    const newComment = ref("");
    const userId = store.userId;

    const { result, refetch } = useQuery(
      GET_ALL_RECIPE_COMMENTS,
      {
        recipe_id: recipeId,
      },
      {
        context: { isPublicQuery: true },
      }
    );

    watch(result, (newResult) => {
      if (newResult && newResult.comments_aggregate) {
        data.value.count = newResult.comments_aggregate.aggregate.count;
        if (newResult.comments_aggregate.nodes) {
          comments.value = newResult.comments_aggregate.nodes;
        }
      }
    });
    const { mutate } = useMutation(ADD_COMMENTS_MUTATION);

    const postComment = async () => {
      if (newComment.value.trim() === "") return;

      try {
        await mutate({
          recipe_id: recipeId,
          user_id: userId,
          comment_text: newComment.value,
        });
        newComment.value = "";
        await refetch();
      } catch (error) {
        console.error("Error posting comment:", error);
      }
    };

    // Format date for displaying comments
    const formatDate = (dateString) => {
      const options = {
        year: "numeric",
        month: "long",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      };
      return new Date(dateString).toLocaleDateString(undefined, options);
    };

    return {
      data,
      comments,
      newComment,
      postComment,
      formatDate,
      userId,
    };
  },
};
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
