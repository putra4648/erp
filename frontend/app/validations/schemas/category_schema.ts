import Joi from "joi";
import type { Category } from "~/types/models/product";

export const CategorySchema = Joi.object<Category>({
  id: Joi.string().allow(""),
  name: Joi.string().required(),
});
