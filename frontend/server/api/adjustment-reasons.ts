import type { AdjustmentReason } from "~/types/models/stock_adjustment";

export default defineEventHandler(async (event) => {
  const method = getMethod(event);

  if (method === "POST") {
    const body = await readBody(event);
    const result = await callBackend<AdjustmentReason>(
      event,
      "/api/adjustment-reason",
      {
        method: "POST",
        body,
      },
    );
    return result;
  }

  const result = await callBackend<AdjustmentReason[]>(
    event,
    "/api/adjustment-reason",
    {
      method: "GET",
    },
  );
  return result;
});
