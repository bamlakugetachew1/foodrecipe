<template>
  <div class="spinner-wrapper">
    <div class="spinner"></div>
  </div>
</template>

<script>
export default {
  name: "Spinner",
};
</script>

<style scoped>
.spinner-wrapper {
  display: flex;
  justify-content: center;
  align-items: flex-start; /* Move the spinner to the top */
  height: 100vh;
  padding-top: 20px; /* Adjust this to control the spacing from the top */
}

.spinner {
  border: 8px solid rgba(0, 0, 0, 0.1);
  border-top: 8px solid #3498db;
  border-radius: 50%;
  width: 60px;
  height: 60px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
