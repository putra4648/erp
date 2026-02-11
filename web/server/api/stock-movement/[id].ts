import type { StockMovement } from "~/types/models/stock_movement";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = getMethod(event);

  if (method === "DELETE") {
    const result = await callBackend(event, `/api/stock-movement/${id}`, {
      method: "DELETE",
    });
    return result;
  }

  if (method === "PUT") {
    const body = await readBody(event);
    const result = await callBackend<StockMovement>(
      event,
      `/api/stock-movement/${id}`,
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
      `/api/stock-movement/${id}`,
      {
        method: "GET",
      },
    );
    return result;
  }
});
