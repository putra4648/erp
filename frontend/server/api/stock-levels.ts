import type { StockLevelResponse } from "~/types/models/stock_level";
import type PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = event.method;

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<StockLevelResponse>(
      event,
      "/api/stock-levels",
      {
        method: "POST",
        body,
      },
    );
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<StockLevelResponse>>(
    event,
    "/api/stock-levels",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
