import { NuxtAuthHandler } from "#auth";
import Auth0Provider from "next-auth/providers/auth0";

export default NuxtAuthHandler({
  secret: useRuntimeConfig().auth0.sessionSecret,
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    Auth0Provider.default({
      clientId: useRuntimeConfig().auth0.clientId,
      clientSecret: useRuntimeConfig().auth0.clientSecret,
      issuer: "https://" + useRuntimeConfig().auth0.domain,
      authorization: {
        params: {
          scope: "openid profile email",
          // POIN PENTING: Tambahkan audience di sini
          audience: useRuntimeConfig().auth0.audience,
        },
      },
    }),
  ],
  callbacks: {
    jwt: ({ token, account }) => {
      if (account) {
        token.accessToken = account.access_token;
        token.idToken = account.id_token;
      }
      return token;
    },
    session: ({ session, token }) => {
      if (session.user) {
        (session as any).accessToken = token.accessToken;
        (session as any).user.id = token.sub;
      }
      return session;
    },
  },
});
