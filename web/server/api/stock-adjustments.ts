import type { StockAdjustment } from "~/types/models/stock_adjustment";
import PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<StockAdjustment>(
      event,
      "/api/stock-adjustment",
      {
        method: "POST",
        body,
      },
    );
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<StockAdjustment>>(
    event,
    "/api/stock-adjustment",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
