// file: ~/server/middleware/auth.ts
import { getServerSession } from "#auth";

export default eventHandler(async (event) => {
  // Skip auth routes to avoid circular checks or unnecessary processing
  if (event.path.startsWith("/api/auth")) {
    return;
  }

  const session = await getServerSession(event);

  // Debug log to help identify the state
  console.log(
    `[Auth Middleware] Path: ${event.path}, Session: ${session ? "Active" : "NULL"}`,
  );

  // If you need to debug cookies, uncomment:
  // console.log('[Auth Middleware] Cookies:', parseCookies(event));
});
