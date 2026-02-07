import Joi from "joi";
import type { Product, UOM } from "~/types/models/product";

export const ProductSchema = Joi.object<Product>({
  id: Joi.string().allow(""),
  sku: Joi.string().required(),
  name: Joi.string().required(),
  min_stock: Joi.number().required(),
  categories: Joi.object<Product["categories"]>({
    id: Joi.string().required(),
    name: Joi.string().required(),
  }).required(),
  uoms: Joi.object<UOM>({
    id: Joi.string().required(),
    name: Joi.string().required(),
  }).required(),
});
