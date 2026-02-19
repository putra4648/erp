import type { Supplier } from "./supplier";
import type { Audit } from "./audit";

export interface Product extends Audit {
  id: string;
  sku: string;
  name: string;
  supplier_id: string;
  supplier?: Supplier;
  categories: Category[];
  uoms: UOM[];
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
