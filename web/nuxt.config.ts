// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: ['@nuxt/ui', '@sidebase/nuxt-auth', "@nuxt/eslint"],
  css: ['~/assets/css/main.css'],
  auth: {
    provider: {
      type: "authjs",
      trustHost: false,
      defaultProvider: "keycloak",
      addDefaultCallbackUrl: true,
    },
    globalAppMiddleware: true,
  },
   eslint: {
    config: {
      stylistic: true // <---
    }
  },
  runtimeConfig: {
    app: {},
    public: {
      serverUrl: "http://localhost:9000",
      clientUrl: "http://localhost:3000",
    },
  },
})