import NextAuth, { DefaultSession } from "next-auth";

declare module "next-auth" {
  interface Session {
    idToken?: string;
    accessToken?: string;
    refreshToken?: string;
    user: {} & DefaultSession["user"];
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    idToken?: string;
    accessToken?: string;
    refreshToken?: string;
  }
}
