export interface StockLevelResponse {
  id: string;
  product_id: string;
  product_name: string;
  warehouse_id: string;
  warehouse_name: string;
  quantity: number;
  last_updated: string;
}

export interface StockLevelRequest {
  product_id?: string;
  warehouse_id?: string;
  search?: string;
  page: number;
  size: number;
}
