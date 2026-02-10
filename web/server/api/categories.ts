import { getServerSession } from "#auth";
import type { Category } from "~/types/models/product";
import PaginationResponse from "../utils/pagination_response";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<Category>(event, "/api/categories", {
      method: "POST",
      body,
    });
    return result;
  }

  const query = getQuery(event);
  const result = await callBackend<PaginationResponse<Category>>(
    event,
    "/api/categories",
    {
      method: "GET",
      query: query,
    },
  );
  return result;
});
