export default defineOAuthAuth0EventHandler({
  config: {
    scope: ["openid", "profile", "email", "offline_access"],
  },
  async onSuccess(event, { user, tokens }) {
    await setUserSession(event, {
      user: {
        id: user.sub,
        name: user.nickname,
        email: user.email,
        image: user.picture,
      },
      secure: {
        access_token: tokens.access_token,
      },
    });
    return sendRedirect(event, "/");
  },
  onError(event, error) {
    console.error("Auth0 Login Error:", error);
    return sendRedirect(event, "/");
  },
});
