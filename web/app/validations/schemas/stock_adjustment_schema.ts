import Joi from "joi";
import { AdjustmentType } from "~/types/enums/adjustment_enum";
import type { StockAdjustment } from "~/types/models/stock_adjustment";
import { Status } from "~/types/enums/status_enum";

export const StockAdjustmentSchema = Joi.object<StockAdjustment>({
  id: Joi.string().optional(),
  type: Joi.string()
    .valid(...Object.values(AdjustmentType))
    .required(),
  reason: Joi.string().required(),
  adjustment_date: Joi.date().required(),
  product_id: Joi.string().required(),
  quantity: Joi.number().integer().min(1).required(),
  notes: Joi.string().optional(),
  status: Joi.string()
    .valid(...Object.values(Status))
    .default(Status.DRAFT),
});
