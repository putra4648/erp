import type { StockAdjustment } from "~/types/models/stock_adjustment";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  const method = getMethod(event);

  if (method === "GET") {
    const result = await callBackend<StockAdjustment>(
      event,
      `/api/stock-adjustment/${id}`,
      {
        method: "GET",
      },
    );
    return result;
  }
});
