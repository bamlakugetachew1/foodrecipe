<template>
  <div class="max-w-full md:max-w-2xl mx-auto mt-7 relative px-4 rounded-full">
    <div class="max-w-full mx-auto mt-4">
      <div class="flex flex-col gap-2 font-sans">
        <label
          for="recipeName"
          class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
        >
          Recipe name
        </label>
        <input
          type="text"
          v-model="recipeName"
          placeholder="Recipe name"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          required
        />

        <label
          for="description"
          class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
        >
          Description
        </label>
        <textarea
          v-model="description"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          placeholder="Write description about the recipe"
          required
        ></textarea>

        <label
          for="preparationTime"
          class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
        >
          Preparation Time In Minutes
        </label>
        <input
          type="text"
          v-model="preparationTime"
          placeholder="Preparation Time"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          required
        />

        <label
          for="category"
          class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
        >
          Select a category
        </label>
        <select
          id="category"
          v-model="selectedCategory"
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
        >
          <option disabled value="">Choose a food category</option>
          <option value="injera">Injera</option>
          <option value="wats">Wats (Stews)</option>
          <option value="vegetarian">Vegetarian Dishes</option>
          <option value="salads">Salads</option>
          <option value="spices">Spices and Condiments</option>
          <option value="grilled_meats">Grilled Meats</option>
          <option value="desserts">Desserts</option>
          <option value="beverages">Beverages</option>
          <option value="breakfast">Breakfast Dishes</option>
        </select>

        <ImageUploader @update-images="handleImageUpdate" />
        <DynamicIngredient @update-ingredients="handleIngredientsUpdate" />
        <DynamicSteps @update-steps="handleStepsUpdate" />

        <button
          class="border p-2 mt-2 bg-blue-500 text-white"
          @click="uploadTexts"
        >
          Add Recipe
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import ImageUploader from "./ImageUploader.vue";
import DynamicIngredient from "./DynamicIngredient.vue";
import DynamicSteps from "./DynamicSteps.vue";
import axios from "axios";
import { useMutation } from "@vue/apollo-composable";
import {
  INSERT_RECIPES_ONE_MUTAION,
  INSERT_INGREDIENTS_MUTATION,
  INSERT_STEPS_MUTATION,
  INSERT_IMAGES_MUTATION,
} from "@/graphql/mutations/creator";
import { useUserStore } from "~/store/userStore";
export default {
  name: "AddUpdateRecipeForm",
  components: {
    ImageUploader,
    DynamicIngredient,
    DynamicSteps,
  },
  setup() {
    const store = useUserStore();
    return {
      recipeName: "",
      description: "",
      preparationTime: "",
      selectedCategory: "",
      parentIngredients: [],
      parentSteps: [],
      parentImages: [],
      parentLocalFiles: [],
      processedImages: [],
      id: store.userId,
      token: store.token,
    };
  },
  methods: {
    async uploadTexts() {
      try {
        const variables = {
          title: this.recipeName,
          description: this.description,
          preparation_time: this.preparationTime,
          category: this.selectedCategory,
          user_id: this.id,
        };
        const { mutate } = useMutation(INSERT_RECIPES_ONE_MUTAION);
        const response = await mutate(variables);
        await this.uploadSteps(response.data.insert_recipes_one.id);
        await this.uploadIngredients(response.data.insert_recipes_one.id);
        this.uploadLocalFiles(response.data.insert_recipes_one.id);
        useRouter().push({
          name: "curator-id",
          params: { id: this.id },
        });
      } catch (error) {
        console.error("Error adding recipe:", error);
      }
    },
    async uploadSteps(recipe_id) {
      try {
        let steps = [];
        for (let i = 0; i < this.parentSteps.length; i++) {
          let data = {
            recipe_id: recipe_id,
            instruction: this.parentSteps[i],
          };
          steps.push(data);
        }
        const { mutate } = useMutation(INSERT_STEPS_MUTATION);
        await mutate({ objects: steps });
      } catch (error) {
        console.error("Error adding steps:", error);
      }
    },
    async uploadIngredients(recipe_id) {
      try {
        let ingredients = [];
        for (let i = 0; i < this.parentIngredients.length; i++) {
          let data = {
            ingredient_name: this.parentIngredients[i].ingredient_name,
            recipe_id: recipe_id,
          };
          ingredients.push(data);
        }
        const { mutate } = useMutation(INSERT_INGREDIENTS_MUTATION);
        await mutate({ objects: ingredients });
      } catch (error) {
        console.error("Error adding ingredients:", error);
      }
    },
    handleIngredientsUpdate(ingredients) {
      this.parentIngredients = ingredients;
    },
    handleStepsUpdate(steps) {
      this.parentSteps = steps;
    },
    handleImageUpdate(imageStore, localFiles) {
      this.parentImages = imageStore;
      this.parentLocalFiles = localFiles;
    },
    uploadLocalFiles(recipe_id) {
      let store = [];
      const formData = new FormData();

      // Prepare form data for local files
      this.parentLocalFiles.forEach((fileData) => {
        formData.append("files", fileData.file);
      });

      // Send files to the backend
      axios
        .post("http://localhost:8081/files/upload", formData, {
          headers: {
            Authorization: `Bearer ${this.token}`,
            "Content-Type": "multipart/form-data",
          },
        })
        .then(async (response) => {
          // Process uploaded file names and create objects for each
          const uploadedFiles = response.data.fileNames.map(
            (fileName, index) => {
              return {
                image_url: `http://localhost:8081/uploads/${fileName}`,
                is_thumbnail: this.parentLocalFiles[index].is_thumbnail, // Assigning thumbnail property
                recipe_id: recipe_id, // Set recipe_id here
              };
            }
          );

          // Merge uploaded files with existing parentImages
          store = [
            ...uploadedFiles.map((file) => ({ ...file })), // Ensure uploaded files are plain objects
            ...this.parentImages.map((img) => ({
              image_url: img.image_url, // Ensure URL is retained
              is_thumbnail: img.is_thumbnail || false, // Use existing thumbnail or default to false
              recipe_id: recipe_id, // Keep existing recipe_id or assign 1 if not available
            })), // Ensure parentImages are also plain
          ];

          // Ensure the first thumbnail image is at the top
          const thumbnailIndex = store.findIndex((img) => img.is_thumbnail);
          if (thumbnailIndex > -1) {
            [store[0], store[thumbnailIndex]] = [
              store[thumbnailIndex],
              store[0],
            ];
          }
          const { mutate } = useMutation(INSERT_IMAGES_MUTATION);
          await mutate({ objects: store });
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>

<style scoped>
/* Add any custom styles if needed */
</style>
