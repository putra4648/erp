import Joi from "joi";
import type { Supplier } from "~/types/models/supplier";

export const SupplierSchema = Joi.object<Supplier>({
  id: Joi.string().required(),
  name: Joi.string().required(),
  contact_person: Joi.string().required(),
  email: Joi.string().email(),
  address: Joi.string().required(),
  is_active: Joi.boolean().required().default(true),
});
