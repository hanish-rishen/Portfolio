// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  modules: ['@nuxtjs/tailwindcss'],
  build: {
    transpile: ['gsap'],
  },
  tailwindcss: {
    config: {
      theme: {
        extend: {
          fontFamily: {
            'cabinet-grotesk': ['Cabinet Grotesk', 'sans-serif'],
            'satoshi': ['Satoshi', 'sans-serif'],
          },
        },
      },
    },
  },
})
