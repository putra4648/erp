export interface StockAdjustment {
  id: string;
  adjustment_no: string;
  warehouse_id: string;
  transaction_date: string;
  status: string;
  note: string;
  created_by?: string;
  approved_by?: string;
  items: StockAdjustmentItem[];
}

export interface StockAdjustmentItem {
  id: string;
  product_id: string;
  product_name?: string;
  reason_id: string;
  reason_name?: string;
  actual_qty: number;
  system_qty: number;
  adjustment_qty: number;
}

export interface AdjustmentReason {
  id: string;
  name: string;
  account_code: string;
}
