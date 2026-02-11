import type { Supplier } from "~/types/models/supplier";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = getMethod(event);

  if (method === "DELETE") {
    const result = await callBackend(event, `/api/supplier/${id}`, {
      method: "DELETE",
    });
    return result;
  }

  if (method === "PUT") {
    const body = await readBody(event);
    const result = await callBackend<Supplier>(event, `/api/supplier/${id}`, {
      method: "PUT",
      body,
    });
    return result;
  }

  if (method === "GET") {
    const result = await callBackend<Supplier>(event, `/api/supplier/${id}`, {
      method: "GET",
    });
    return result;
  }
});
