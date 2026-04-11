import type { NuxtPage } from "nuxt/schema";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxt/eslint", "nuxt-auth-utils"],
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
  runtimeConfig: {
    public: {
      serverUrl: process.env.NUXT_SERVER_URL,
    },
  },
  hooks: {
    "pages:extend"(pages: NuxtPage[]) {
      function setMiddleware(pages: NuxtPage[]) {
        for (const page of pages) {
          if (page.path !== "/") {
            page.meta ||= {};
            const mw = page.meta.middleware;
            if (!mw) {
              page.meta.middleware = ["auth"];
            } else if (Array.isArray(mw)) {
              if (!mw.includes("auth")) mw.push("auth");
            } else if (mw !== "auth") {
              page.meta.middleware = [mw, "auth"];
            }
          }

          if (page.children) {
            setMiddleware(page.children);
          }
        }
      }
      setMiddleware(pages);
    },
  },
});
