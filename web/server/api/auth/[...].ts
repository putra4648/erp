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
      // 1. Login Pertama Kali
      if (account && profile) {
        console.log(account.expires_at);
        return {
          accessToken: account.access_token,
          accessTokenExpires: Date.now() + (account.expires_at ?? 0) * 1000,
          refreshToken: account.refresh_token,
          idToken: account.id_token,
          roles: (profile as any).realm_access?.roles || [],
          groups: (profile as any).groups || [],
          user: token,
        };
      }

      // 2. Belum Kadaluarsa (Kita beri buffer 10 detik)
      if (Date.now() < (token.accessTokenExpires as number) - 10000) {
        return token;
      }

      // 3. Sudah Kadaluarsa -> Refresh!
      return refreshAccessToken(token);
    },
    async session({ session, token }) {
      session.idToken = token.idToken;
      session.user.roles = token.roles;
      session.user.groups = token.groups;
      session.accessToken = token.accessToken;
      session.refreshToken = token.refreshToken;
      return session;
    },
  },
});

async function refreshAccessToken(token: any) {
  const config = useRuntimeConfig();
  try {
    const url = `${config.public.keycloakUrl}/protocol/openid-connect/token`;
    const response = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: new URLSearchParams({
        client_id: process.env.KEYCLOAK_CLIENT_ID || "",
        client_secret: process.env.KEYCLOAK_CLIENT_SECRET || "",
        grant_type: "refresh_token",
        refresh_token: token.refreshToken,
      }),
    });

    const refreshedTokens = await response.json();

    if (!response.ok) throw refreshedTokens;

    return {
      ...token,
      accessToken: refreshedTokens.access_token,
      accessTokenExpires: Date.now() + refreshedTokens.expires_in * 1000,
      refreshToken: refreshedTokens.refresh_token ?? token.refreshToken, // Pakai yang baru jika ada
    };
  } catch (error) {
    console.error("Error refreshing access token", error);
    return { ...token, error: "RefreshAccessTokenError" };
  }
}
