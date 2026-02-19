import type { StockMovement } from "~/types/models/stock_movement";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = event.method;

  if (method === "DELETE") {
    const result = await callBackend(event, `/api/stock-movements/${id}`, {
      method: "DELETE",
    });
    return result;
  }

  if (method === "PUT") {
    const body = await readBody(event);
    const result = await callBackend<StockMovement>(
      event,
      `/api/stock-movements/${id}`,
      {
        method: "PUT",
        body,
      },
    );
    return result;
  }

  if (method === "GET") {
    const result = await callBackend<StockMovement>(
      event,
      `/api/stock-movements/${id}`,
      {
        method: "GET",
      },
    );
    return result;
  }
});
