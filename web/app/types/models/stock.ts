import type { Status } from "../enums/status_enum";
import type { TransactionType } from "../enums/transaction_enum";
import type { Audit } from "./audit";
import type { Customer } from "./customer";
import type { Product } from "./product";
import type { Supplier } from "./supplier";
import type { Warehouse } from "./warehouse";

export interface Stock {
  id: string;
  transaction_no: string;
  type?: TransactionType;
  source_warehouse: Warehouse;
  target_warehouse: Warehouse;
  supplier: Supplier;
  customer: Customer;
  status: Status;
  transaction_date: Date;
  notes: string;
}

export interface StockDetail {
  id: string;
  stock: Stock;
  product: Product;
  quantity: number;
  unit_cost: number;
  reason_code: string;
}
