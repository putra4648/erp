declare module "#auth-utils" {
  interface User {
    id: string;
    name: string;
    email: string;
    image: string;
  }

  interface UserSession {
    user?: User;
    tokens?: {
      access_token: string;
      refresh_token?: string;
      id_token?: string;
      expires_in?: number;
    };
  }
}

export {};
