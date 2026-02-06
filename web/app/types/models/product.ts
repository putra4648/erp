import type { Audit } from "./audit";

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
