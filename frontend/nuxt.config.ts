// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxt/eslint", "@sidebase/nuxt-auth"],
  css: ["~/assets/css/main.css"],
  eslint: {
    config: {
      stylistic: true, // <---
    },
  },
  app: {
    head: {
      title: "ERP System",
    },
  },
  auth: {
    provider: {
      type: "authjs",
      trustHost: true,
      defaultProvider: "auth0",
      addDefaultCallbackUrl: true,
    },
    globalAppMiddleware: true,
  },
  runtimeConfig: {
    auth0: {
      domain: process.env.NUXT_AUTH0_DOMAIN,
      clientId: process.env.NUXT_AUTH0_CLIENT_ID,
      clientSecret: process.env.NUXT_AUTH0_CLIENT_SECRET,
      sessionSecret: process.env.NUXT_AUTH0_SESSION_SECRET,
      appBaseUrl: process.env.NUXT_AUTH0_APP_BASE_URL,
      audience: process.env.NUXT_AUTH0_AUDIENCE,
    },
    public: {
      serverUrl: process.env.NUXT_SERVER_URL,
      auth0: {
        domain: process.env.NUXT_AUTH0_DOMAIN,
        clientId: process.env.NUXT_AUTH0_CLIENT_ID,
      },
    },
  },
});
