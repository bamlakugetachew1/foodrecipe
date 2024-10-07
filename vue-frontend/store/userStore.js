import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: () => ({
    userId: null,
    token: null,
  }),
  actions: {
    setUserId(userId) {
      this.userId = userId;
      if (process.client) {
        // Check if we are in the client environment
        localStorage.setItem("userid", userId); // Save userId to localStorage
      }
    },
    setToken(token) {
      this.token = token;
      if (process.client) {
        // Check if we are in the client environment
        localStorage.setItem("token", token); // Save token to localStorage
      }
    },
    logout() {
      this.userId = null;
      this.token = null;
      if (process.client) {
        // Check if we are in the client environment
        localStorage.removeItem("userid"); // Remove userId from localStorage
        localStorage.removeItem("token"); // Remove token from localStorage
      }
    },
  },
  persist: process.client
    ? {
        enabled: true, // Enable persistence
        strategies: [
          {
            key: "user", // The key in localStorage
            storage: localStorage, // Use localStorage
          },
        ],
      }
    : false, // Disable persistence on the server
});
