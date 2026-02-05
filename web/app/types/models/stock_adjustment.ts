import type { AdjustmentType } from "../enums/adjustment_enum";
import type { Status } from "../enums/status_enum";

export type StockAdjustment = {
  id: string;
  type: AdjustmentType;
  reason: string;
  adjustment_date: string | undefined;
  product_id: string;
  quantity: number;
  notes?: string;
  status: Status;
};
