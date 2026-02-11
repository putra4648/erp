import Joi from "joi";
import { TransactionType } from "~/types/enums/transaction_enum";
import type { Customer } from "~/types/models/customer";
import type { Stock } from "~/types/models/stock";
import type { Supplier } from "~/types/models/supplier";
import type { Warehouse } from "~/types/models/warehouse";

export const StockSchema = Joi.object<Stock>({
  id: Joi.string().allow(""),
  type: Joi.string()
    .valid(...Object.keys(TransactionType).filter((key) => isNaN(Number(key))))
    .required(),
  transaction_no: Joi.string().allow(""),
  notes: Joi.string().allow(""),
  source_warehouse: Joi.object<Warehouse>({
    id: Joi.string().empty("").optional(),
    name: Joi.string().empty("").optional(),
  })
    .allow(null, {})
    .optional()
    .unknown(true),
  target_warehouse: Joi.object<Warehouse>({
    id: Joi.string().empty("").optional(),
    name: Joi.string().empty("").optional(),
  })
    .allow(null, {})
    .optional()
    .unknown(true),
  status: Joi.string().required(),
  supplier: Joi.object<Supplier>({
    id: Joi.string().empty("").optional(),
    name: Joi.string().empty("").optional(),
  })
    .allow(null, {})
    .optional()
    .unknown(true),
  customer: Joi.object<Customer>({
    id: Joi.string().empty("").optional(),
    name: Joi.string().empty("").optional(),
  })
    .allow(null, {})
    .optional()
    .unknown(true),
  transaction_date: Joi.date().required(),
});
