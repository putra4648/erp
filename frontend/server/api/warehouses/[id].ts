import type { Warehouse } from "~/types/models/warehouse";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = getMethod(event);

  if (method === "DELETE") {
    const result = await callBackend(event, `/api/warehouse/${id}`, {
      method: "DELETE",
    });
    return result;
  }

  if (method === "PUT") {
    const body = await readBody(event);
    const result = await callBackend<Warehouse>(event, `/api/warehouse/${id}`, {
      method: "PUT",
      body,
    });
    return result;
  }

  if (method === "GET") {
    const result = await callBackend<Warehouse>(event, `/api/warehouse/${id}`, {
      method: "GET",
    });
    return result;
  }
});
