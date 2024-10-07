<script>
export default {
  name: "RecipeCard",
  props: {
    id: String,
    title: String,
    category: String,
    preparation_time: String,
    description: String,
    isAdmin: Boolean,
    url: String,
  },
  methods: {
    handleEdit() {
      this.$emit("edit", this.id);
    },
    handleDelete() {
      this.$emit("delete", this.id);
    },
  },
};
</script>

<template>
  <div
    class="max-w-sm mb-2 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700"
  >
    <a href="#">
      <img class="rounded-t-lg" :src="url" alt="Image description" />
    </a>
    <div class="p-5">
      <div class="flex justify-between">
        <p class="font-serif text-lg text-gray-600">{{ title }}</p>
        <p class="font-serif text-lg text-gray-600">{{ category }}</p>
        <div class="flex gap-3">
          <i class="fa fa-clock mt-2" aria-hidden="true" id="space"></i>

          <p class="font-serif text-lg text-gray-600">{{ preparation_time }}</p>
        </div>
      </div>
      <div class="flex justify-between mt-1">
        <button v-if="isAdmin" @click="handleDelete">
          <i class="fa fa-trash" aria-hidden="true" id="space"></i>
        </button>
        <a>
          <nuxt-link
            :to="{
              name: 'recipes-id',
              params: { id: id },
              query: {
                title: title,
                category: category,
                preparation_time: preparation_time,
                description: description,
              },
            }"
          >
            <i class="fa fa-arrow-right" aria-hidden="true" id="space"></i>
          </nuxt-link>
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
