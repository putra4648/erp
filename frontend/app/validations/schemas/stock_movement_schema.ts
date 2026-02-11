import Joi from "joi";
import type { StockMovementDTO } from "~/types/models/stock_movement";

export const StockMovementSchema = Joi.object<StockMovementDTO>({
  id: Joi.string().allow(""),
  type: Joi.string().valid("IN", "OUT", "TRANSFER").required(),
  origin_warehouse_id: Joi.string().when("type", {
    is: Joi.valid("OUT", "TRANSFER"),
    then: Joi.required(),
    otherwise: Joi.string().allow(null, ""),
  }),
  destination_warehouse_id: Joi.string().when("type", {
    is: Joi.valid("IN", "TRANSFER"),
    then: Joi.required(),
    otherwise: Joi.string().allow(null, ""),
  }),
  reference_no: Joi.string().allow(""),
  status: Joi.string().required(),
  transaction_date: Joi.string().required(),
  note: Joi.string().allow(""),
  items: Joi.array()
    .items(
      Joi.object({
        id: Joi.string().allow(""),
        product_id: Joi.string().required(),
        quantity: Joi.number().greater(0).required(),
        note: Joi.string().allow(""),
      }),
    )
    .min(1)
    .required(),
});
