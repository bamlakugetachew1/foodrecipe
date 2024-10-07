<template>
  <div
    class="mb-4 block space-y-2 md:space-y-0 md:flex items-center gap-4 mt-2 md:mt-4 ml-2 mr-2 lg:mr-0 lg:ml-4"
    id="shift"
  >
    <select
      @change="onFilterChange"
      class="border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 transition w-full sm:w-auto"
    >
      <option value="all">ALL Foods</option>
      <option value="time">Filter by time</option>
      <option value="title">Filter by title</option>
    </select>

    <div
      class="flex items-center border border-gray-300 rounded-lg overflow-hidden"
    >
      <input
        v-model="inputData"
        class="px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 w-full"
        type="text"
        placeholder="Search..."
      />
      <button
        :disabled="inputData.length < 1"
        @click="searchRecipes"
        class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 transition focus:outline-none"
      >
        <svg
          class="h-5 w-5"
          fill="currentColor"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
        >
          <path
            d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z"
          />
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "RecipeFilter",
  data() {
    return {
      inputData: "",
      selectedFilter: "all",
    };
  },
  methods: {
    onFilterChange(event) {
      this.selectedFilter = event.target.value;
      this.$emit("filter-change", this.selectedFilter);
    },
    searchRecipes() {
      this.$router.push({
        name: "recipes-search",
        query: { searchQuery: this.inputData },
      });
    },
  },
};
</script>

<style scoped></style>
