import { ROUTE_PERMISSIONS } from "~/constants/route_permission";

export default defineNuxtRouteMiddleware((to) => {
  const { canAccess } = useNavAccess();

  // Skip jika login atau unauthorized agar tidak infinite loop
  if (["/login", "/unauthorized"].includes(to.path)) return;

  // Cari config rute berdasarkan wildcard
  const activeConfig = ROUTE_PERMISSIONS.find((config) => {
    const pattern = config.path.replace(/\//g, "\\/").replace(/\*\*/g, ".*");
    return new RegExp(`^${pattern}$`).test(to.path);
  });

  // LOGIKA:
  // Jika rute ini terdaftar (activeConfig)
  // DAN canAccess mengembalikan false (user tidak berhak)
  if (activeConfig && !canAccess(activeConfig)) {
    console.warn("Akses ditolak ke:", to.path);
    return navigateTo("/unauthorized");
  }
});
