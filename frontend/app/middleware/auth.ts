export default defineNuxtRouteMiddleware((to, from) => {
  const { user } = useUserSession();

  if (!user.value) {
    window.location.href = "/auth/auth0";
    return;
  }
});
