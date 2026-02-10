import type { Supplier } from "~/types/models/supplier";
import PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<Supplier>(event, "/api/supplier", {
      method: "POST",
      body,
    });
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<Supplier>>(
    event,
    "/api/supplier",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
