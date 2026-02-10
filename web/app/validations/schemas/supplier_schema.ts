import Joi from "joi";
import type { Supplier } from "~/types/models/supplier";

export const SupplierSchema = Joi.object<Supplier>({
  id: Joi.string().allow(""),
  name: Joi.string().required(),
  code: Joi.string().required(),
  email: Joi.string().email().allow(""),
  phone: Joi.string().allow(""),
  address: Joi.string().allow(""),
  is_active: Joi.boolean().default(true),
});
