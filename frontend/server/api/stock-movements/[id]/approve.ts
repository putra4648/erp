import { StockMovement } from "~/types/models/stock_movement";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = event.method;
  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<StockMovement>(
      event,
      `/api/stock-movements/${id}/approve`,
      {
        method: "POST",
        body,
      },
    );
    return result;
  }
});
