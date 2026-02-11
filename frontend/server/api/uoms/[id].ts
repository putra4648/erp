import type { UOM } from "~/types/models/product";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = getMethod(event);

  if (method === "DELETE") {
    const result = await callBackend(event, `/api/uoms/${id}`, {
      method: "DELETE",
    });
    return result;
  }

  if (method === "PUT") {
    const body = await readBody(event);
    const result = await callBackend<UOM>(event, `/api/uoms/${id}`, {
      method: "PUT",
      body,
    });
    return result;
  }

  if (method === "GET") {
    const result = await callBackend<UOM>(event, `/api/uoms/${id}`, {
      method: "GET",
    });
    return result;
  }
});
