// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: false },
  modules: ["@nuxtjs/tailwindcss"],
  vite: {
    optimizeDeps: {
      include: ['@apollo/client', 'graphql'],
    },
  },
  plugins: ['~/plugins/apollo-client.js','~/plugins/fontawesome.js','~/plugins/piniaPersistedState.js'],
  css: [
    'slick-carousel/slick/slick.css',
    'slick-carousel/slick/slick-theme.css'
  ],
});
