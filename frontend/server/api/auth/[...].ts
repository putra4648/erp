import { NuxtAuthHandler } from "#auth";
import KeycloakProvider from "next-auth/providers/keycloak";

export default NuxtAuthHandler({
  secret: "some random string",
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    KeycloakProvider.default({
      issuer: useRuntimeConfig().keycloakIssuer,
      clientId: useRuntimeConfig().public.clientId || "",
      clientSecret: useRuntimeConfig().keycloakClientSecret || "",
      authorization: {
        params: {
          scope: "openid profile email", // Pastikan ini ada
        },
      },
    }),
  ],
  callbacks: {
    async jwt({ token, account, profile, user }) {
      // 1. Initial login
      if (account) {
        console.log("Initial Login - Account values:", {
          expires_at: account.expires_at,
          refresh_expires_in: account.refresh_expires_in,
        });

        token.accessToken = account.access_token;
        token.accessTokenExpires = (account.expires_at as number) * 1000;
        token.refreshToken = account.refresh_token;
        token.refreshTokenExpires =
          Date.now() + (account.refresh_expires_in as number) * 1000;
        token.roles = (profile as any).realm_access?.roles || [];
        token.groups = (profile as any).groups || [];

        console.log("Initial JWT token state:", {
          accessTokenExpires: new Date(
            token.accessTokenExpires as number,
          ).toLocaleString(),
          now: new Date().toLocaleString(),
        });

        return token; // Return immediately after setting initial token
      }

      // 2. Refresh token before 5 minutes of expiration (300,000 ms)
      const now = Date.now();
      const expiresAt = token.accessTokenExpires as number;

      // Jika sudah ada error sebelumnya atau masih belum waktunya refresh, jangan refresh
      if (token.error || now < expiresAt - 300000) {
        return token;
      }

      console.log(
        "Token near expiration, attempting refresh. Now:",
        new Date(now).toLocaleString(),
        "Expires:",
        new Date(expiresAt).toLocaleString(),
      );

      // 3. Expired or near expiration -> Refresh!
      if (!token.refreshToken) {
        console.error("Missing refresh token in token object");
        return { ...token, error: "MissingRefreshToken" };
      }
      return refreshAccessToken(token);
    },
    session({ session, token }) {
      return {
        ...session,
        user: {
          ...session.user,
          roles: token.roles,
          groups: token.groups,
        },
        accessToken: token.accessToken,
        refreshToken: token.refreshToken,
      };
    },
  },
});

async function refreshAccessToken(token: any) {
  const config = useRuntimeConfig();
  try {
    const url = `${config.public.keycloakUrl}/realms/erp/protocol/openid-connect/token`;

    console.log("Refreshing token with URL:", url);
    console.log("Client ID:", config.public.clientId);

    const response = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: new URLSearchParams({
        client_id: config.public.clientId || "",
        client_secret: config.keycloakClientSecret || "",
        grant_type: "refresh_token",
        refresh_token: token.refreshToken,
      }),
    });

    const refreshedTokens = await response.json();

    if (!response.ok) {
      console.error("Keycloak refresh error response:", refreshedTokens);
      throw refreshedTokens;
    }

    console.log("Successfully refreshed tokens");

    return {
      ...token,
      accessToken: refreshedTokens.access_token,
      accessTokenExpires: Date.now() + refreshedTokens.expires_in * 1000,
      refreshToken: refreshedTokens.refresh_token ?? token.refreshToken, // Pakai yang baru jika ada
      refreshTokenExpires: refreshedTokens.refresh_expires_in
        ? Date.now() + refreshedTokens.refresh_expires_in * 1000
        : token.refreshTokenExpires,
    };
  } catch (error) {
    console.error("Error refreshing access token details:", error);
    return { ...token, error: "RefreshAccessTokenError" };
  }
}
