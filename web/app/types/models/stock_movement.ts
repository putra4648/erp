import type { Status } from "../enums/status_enum";
import type { Warehouse } from "./warehouse";
import type { Product } from "./product";

export interface StockMovement {
  id: string;
  movement_no: string;
  type: string;
  origin_warehouse_id?: string;
  origin_warehouse?: Warehouse;
  destination_warehouse_id?: string;
  destination_warehouse?: Warehouse;
  reference_no: string;
  status: Status;
  transaction_date: string;
  note: string;
  items: StockMovementItem[];
}

export interface StockMovementItem {
  id: string;
  product_id: string;
  product?: Product;
  quantity: number;
  note: string;
}

export interface StockMovementDTO {
  id?: string;
  type: string;
  origin_warehouse_id?: string;
  destination_warehouse_id?: string;
  reference_no?: string;
  status: Status;
  transaction_date?: string;
  note?: string;
  items: StockMovementItemDTO[];
}

export interface StockMovementItemDTO {
  id?: string;
  product_id: string;
  quantity: number;
  note?: string;
}
