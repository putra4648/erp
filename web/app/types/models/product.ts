import type { TransactionType } from "../enums/transaction_enum";
import type { Audit } from "./audit";
import type { Warehouse } from "./warehouse";

export interface Product extends Audit {
  id: string;
  sku: string;
  name: string;
  category: Category;
  uom: UOM;
  min_stock: number;
}

export interface Category {
  id: string;
  name: string;
}

export interface UOM {
  id: string;
  name: string;
}
