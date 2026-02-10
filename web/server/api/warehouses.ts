import type { Warehouse } from "~/types/models/warehouse";
import PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<Warehouse>(event, "/api/warehouse", {
      method: "POST",
      body,
    });
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<Warehouse>>(
    event,
    "/api/warehouse",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
