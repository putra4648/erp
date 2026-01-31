import type { TransactionType } from "../enums/transaction_enum";
import type { Audit } from "./audit";
import type Product from "./product";
import type { Supplier } from "./supplier";
import type { Warehouse } from "./warehouse";

export interface Stock extends Audit {
    id: string;
    product: Product
    warehouse: Warehouse
    supplier: Supplier,
    type: TransactionType,
    quantity: number,
    reference_no: string,
}