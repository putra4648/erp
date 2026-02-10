import Joi from "joi";
import type { Product } from "~/types/models/product";
import type { Warehouse, StockLevel } from "~/types/models/warehouse";

export const WarehouseSchema = Joi.object<Warehouse>({
  id: Joi.string().allow(""),
  name: Joi.string().required(),
  code: Joi.string().required(),
  is_active: Joi.boolean().default(true),
  stock_levels: Joi.array<StockLevel>()
    .items(
      Joi.object({
        id: Joi.string().allow(""),
        product: Joi.object<Product>({
          name: Joi.string().required(),
        }),
        quantity: Joi.number().min(1).required(),
      }),
    )
    .default([]),
});
