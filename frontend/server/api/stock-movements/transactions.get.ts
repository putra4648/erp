import type { StockTransactionResponse } from "~/types/models/stock_transaction";
import type PaginationResponse from "../../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const query = getQuery(event);
  const result = await callBackend<
    PaginationResponse<StockTransactionResponse>
  >(event, "/api/stock-movements/transactions", {
    method: "GET",
    query: query,
  });
  return result;
});
