import type { H3Event } from "h3";
import type { NitroFetchOptions, NitroFetchRequest } from "nitropack";
import { getServerSession } from "#auth";

// Fungsi ini otomatis tersedia di semua file dalam folder server/api
export const callBackend = async <T>(
  event: H3Event,
  url: NitroFetchRequest,
  options: NitroFetchOptions<NitroFetchRequest>,
): Promise<T> => {
  const session = await getServerSession(event);
  const config = useRuntimeConfig(event);
  // Logika auth bisa ditaruh di sini agar terpusat

  return $fetch<T>(url, {
    baseURL: String(config.public.serverUrl),
    ...options,
    onRequest: ({ request, options, error }) => {
      if (session?.accessToken) {
        options.headers.set("Authorization", `Bearer ${session.accessToken}`);
      }
    },
  });
};
