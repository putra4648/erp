export interface StockTransactionResponse {
  id: string;
  product_id: string;
  product_name: string;
  warehouse_id: string;
  warehouse_name: string;
  supplier_id?: string;
  supplier_name?: string;
  type: "IN" | "OUT" | "ADJUST";
  quantity: number;
  reference_no: string;
  created_at: string;
}

export interface StockTransactionRequest {
  product_id?: string;
  warehouse_id?: string;
  page: number;
  size: number;
}
