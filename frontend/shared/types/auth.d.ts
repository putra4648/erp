// shared/types/auth.d.ts
declare module "#auth-utils" {
  interface User {
    // Add your own fields
    id: string;
    name: string;
    email: string;
    image: string;
  }

  interface UserSession {
    // Add your own fields
  }

  interface SecureSessionData {
    // Add your own fields
    access_token: string;
  }
}

export {};
