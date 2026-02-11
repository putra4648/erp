import type { StockAdjustment } from "~/types/models/stock_adjustment";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");

  const result = await callBackend<StockAdjustment>(
    event,
    `/api/stock-adjustment/${id}/void`,
    {
      method: "POST",
    },
  );
  return result;
});
