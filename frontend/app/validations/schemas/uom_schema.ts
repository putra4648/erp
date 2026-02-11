import Joi from "joi";
import type { UOM } from "~/types/models/product";

export const UOMSchema = Joi.object<UOM>({
  id: Joi.string().allow(""),
  name: Joi.string().required(),
});
