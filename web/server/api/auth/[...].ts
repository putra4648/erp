import { NuxtAuthHandler } from "#auth";
import KeycloakProvider from "next-auth/providers/keycloak";

export default NuxtAuthHandler({
  secret: "some random string",
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    KeycloakProvider.default({
      issuer: process.env.KEYCLOAK_ISSUER,
      clientId: process.env.KEYCLOAK_CLIENT_ID || "",
      clientSecret: process.env.KEYCLOAK_CLIENT_SECRET || "",
      authorization: {
        params: {
          scope: "openid profile email", // Pastikan ini ada
        },
      },
    }),
  ],
  callbacks: {
    async jwt({ token, account, profile }) {
      if (account) {
        token.idToken = account.id_token;
        token.roles = (profile as any).realm_access.roles;
        token.groups = (profile as any).groups;
      }
      return token;
    },
    redirect(params) {
      return params.baseUrl;
    },
    async session({ session, token }) {
      session.idToken = token.idToken;
      session.user.roles = token.roles;
      session.user.groups = token.groups;
      return session;
    },
  },
});
