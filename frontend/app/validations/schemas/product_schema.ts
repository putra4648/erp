import Joi from "joi";
import type { Product, UOM } from "~/types/models/product";

export const ProductSchema = Joi.object<Product>({
  id: Joi.string().allow(""),
  sku: Joi.string().required(),
  name: Joi.string().required(),
  min_stock: Joi.number().required(),
  categories: Joi.array().items(Joi.object()).required(),
  uoms: Joi.array().items(Joi.object()).required(),
});
