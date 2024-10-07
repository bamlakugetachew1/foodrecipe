<template>
  <div class="no-recipes-container">
    <div class="message-container">
      <h1 class="title">Oops! No Recipes Found</h1>
      <p class="description">
        We couldn't find any recipes matching your search.
      </p>
      <p class="suggestion">Try searching for something else!</p>
    </div>
    <nuxt-link to="/" class="back-button">Back to Home</nuxt-link>
  </div>
</template>

<script>
export default {
  name: "NoRecipesFound",
};
</script>

<style scoped>
.no-recipes-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  text-align: center;
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.message-container {
  margin-bottom: 20px;
}

.title {
  font-size: 2rem;
  font-weight: bold;
  color: #e74c3c;
}

.description {
  font-size: 1.25rem;
  color: #333;
}

.suggestion {
  font-size: 1rem;
  color: #555;
  margin-top: 10px;
}

.back-button {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #3498db;
  color: white;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.back-button:hover {
  background-color: #2980b9;
}
</style>
