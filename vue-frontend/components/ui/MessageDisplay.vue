<template>
  <h1
    class="font-bold leading-tight tracking-tight text-gray-900 dark:text-white"
  >
    {{ message }}
  </h1>
</template>

<script>
export default {
  name: "MessageDisplay",
  props: {
    message: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
</style>
