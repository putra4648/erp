import type { Audit } from "./audit";
import type { Product } from "./product";

export interface Warehouse {
  id: string;
  name: string;
  code: string;
  is_active: boolean;
  stock_levels: StockLevel[];
}

export interface StockLevel extends Audit {
  id: string;
  product: Product;
  warehouse?: Warehouse;
  quantity: number;
}
