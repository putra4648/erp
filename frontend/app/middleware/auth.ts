export default defineNuxtRouteMiddleware((to, from) => {
  const { user } = useUserSession();

  if (!user.value) {
    return navigateTo("/auth/auth0", { external: true });
  }
});
