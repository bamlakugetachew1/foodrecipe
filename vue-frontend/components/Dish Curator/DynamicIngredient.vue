<template>
  <div>
    <h2 class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
      Ingredients
    </h2>
    <div class="flex flex-col gap-2">
      <div class="flex gap-2">
        <input
          v-model="newIngredient"
          type="text"
          placeholder="Add an ingredient"
          class="p-1 border w-4/5 border-solid border-gray-300 focus:outline-none rounded-sm"
        />
        <button
          @click="addIngredient"
          class="border p-2 bg-blue-500 text-white"
        >
          Add
        </button>
      </div>
      <div v-if="ingredients.length > 0" class="mt-3">
        <ul>
          <li
            v-for="(ingredient, index) in ingredients"
            :key="index"
            class="flex justify-between items-center border-b py-1"
          >
            <span>{{ ingredient.ingredient_name }}</span>
            <button @click="removeIngredient(index)" class="text-red-500">
              <i class="fa fa-trash"></i>
            </button>
          </li>
        </ul>
      </div>
      <div v-else>
        <p>No ingredients added yet.</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DynamicIngredient",
  data() {
    return {
      newIngredient: "",
      ingredients: [],
    };
  },
  methods: {
    addIngredient() {
      if (this.newIngredient.trim() !== "") {
        this.ingredients.push({ ingredient_name: this.newIngredient.trim() });
        this.$emit("update-ingredients", this.ingredients); // Emit to parent
        this.newIngredient = "";
      }
    },
    removeIngredient(index) {
      this.ingredients.splice(index, 1);
      this.$emit("update-ingredients", this.ingredients); // Emit to parent
    },
  },
};
</script>

<style scoped></style>
