import Joi from "joi";
import type {
  StockAdjustment,
  StockAdjustmentItem,
} from "~/types/models/stock_adjustment";

export const StockAdjustmentSchema = Joi.object<StockAdjustment>({
  id: Joi.string().optional().allow(""),
  warehouse_id: Joi.string().required().messages({
    "any.required": "Warehouse is required",
  }),
  transaction_date: Joi.string().required().messages({
    "any.required": "Transaction date is required",
  }),
  status: Joi.string().required().messages({
    "any.required": "Status is required",
  }),
  adjustment_no: Joi.string().optional().allow(""),
  note: Joi.string().optional().allow(""),
  created_by: Joi.string().optional().allow(""),
  approved_by: Joi.string().optional().allow(""),
  items: Joi.array()
    .items(
      Joi.object<StockAdjustmentItem>({
        id: Joi.string().optional().allow(""),
        product_id: Joi.string().required().messages({
          "any.required": "Product is required",
        }),
        product_name: Joi.string().optional().allow(""),
        reason_id: Joi.string().required().messages({
          "any.required": "Reason is required",
        }),
        reason_name: Joi.string().optional().allow(""),
        actual_qty: Joi.number().required().messages({
          "any.required": "Actual quantity is required",
        }),
        system_qty: Joi.number().required().messages({
          "any.required": "System quantity is required",
        }),
        adjustment_qty: Joi.number().required().messages({
          "any.required": "Adjustment quantity is required",
        }),
      }),
    )
    .min(1)
    .required()
    .messages({
      "array.min": "At least one item is required",
    }),
});

export const AdjustmentReasonSchema = Joi.object({
  id: Joi.string().optional().allow(""),
  name: Joi.string().required(),
  account_code: Joi.string().required(),
});
