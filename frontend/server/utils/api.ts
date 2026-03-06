import type { H3Event } from "h3";
import type { NitroFetchOptions, NitroFetchRequest } from "nitropack";

// Fungsi ini otomatis tersedia di semua file dalam folder server/api
export const callBackend = async <T>(
  event: H3Event,
  url: NitroFetchRequest,
  options: NitroFetchOptions<NitroFetchRequest>,
): Promise<T> => {
  const auth = useAuth0(event);
  const token = await auth.getAccessToken();
  const config = useRuntimeConfig(event);

  console.log(token.accessToken);

  return $fetch<T>(url, {
    baseURL: String(config.public.serverUrl),
    ...options,
    onRequest: ({ request, options, error }) => {
      if (token.accessToken) {
        options.headers.set("Authorization", `Bearer ${token.accessToken}`);
      }
    },
  });
};
