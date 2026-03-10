import type { Product } from "~/types/models/product";
import PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<Product>(event, "/api/products", {
      method: "POST",
      body,
    });
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<Product>>(
    event,
    "/api/products",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
