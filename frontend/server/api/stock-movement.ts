import type { StockMovement } from "~/types/models/stock_movement";
import type PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<StockMovement>(
      event,
      "/api/stock-movement",
      {
        method: "POST",
        body,
      },
    );
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<StockMovement>>(
    event,
    "/api/stock-movement",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
