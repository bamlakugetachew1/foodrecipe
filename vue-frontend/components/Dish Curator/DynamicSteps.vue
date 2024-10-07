<template>
  <div>
    <h2 class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
      Steps
    </h2>
    <div class="flex flex-col gap-2">
      <div class="flex gap-2">
        <input
          v-model="newStep"
          type="text"
          placeholder="Add a step"
          class="p-1 border w-4/5 border-solid border-gray-300 focus:outline-none rounded-sm"
        />
        <button @click="addStep" class="border p-2 bg-blue-500 text-white">
          Add
        </button>
      </div>
      <div v-if="steps.length > 0" class="mt-3">
        <ul>
          <li
            v-for="(step, index) in steps"
            :key="index"
            class="flex justify-between items-center border-b py-1"
          >
            <span>{{ step }}</span>
            <button @click="removeStep(index)" class="text-red-500">
              <i class="fa fa-trash"></i>
            </button>
          </li>
        </ul>
      </div>
      <div v-else>
        <p>No steps added yet.</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DynamicSteps",
  data() {
    return {
      newStep: "",
      steps: [],
    };
  },
  methods: {
    addStep() {
      if (this.newStep.trim() !== "") {
        this.steps.push(this.newStep.trim());
        this.$emit("update-steps", this.steps);
        this.newStep = "";
      }
    },
    removeStep(index) {
      this.steps.splice(index, 1);
      this.$emit("update-steps", this.steps);
    },
  },
};
</script>

<style scoped></style>
