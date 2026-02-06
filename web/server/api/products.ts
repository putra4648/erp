import { getServerSession } from "#auth";
import type { Product } from "~/types/models/product";

export default defineEventHandler(async (event) => {
  const result = await callBackend<Product[]>(event, "/api/products", {
    method: "GET",
  });
  return result || [];
});
