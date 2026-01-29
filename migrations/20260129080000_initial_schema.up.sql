-- EKSTENSI UNTUK UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ==========================================
-- FASE 1: APPROVAL SYSTEM
-- ==========================================
CREATE TABLE approval_workflows (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    doc_code VARCHAR(50) UNIQUE NOT NULL, -- Contoh: 'STOCK_ADJ'
    doc_name VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE approval_steps (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workflow_id UUID REFERENCES approval_workflows(id),
    step_order INT NOT NULL,
    target_group_name VARCHAR(100) NOT NULL, -- Nama Group dari Keycloak
    min_approver INT DEFAULT 1
);

CREATE TABLE approval_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workflow_id UUID REFERENCES approval_workflows(id),
    reference_id UUID NOT NULL, -- ID dari tabel Bisnis (misal: stock_adjustment_id)
    current_step INT DEFAULT 1,
    status VARCHAR(20) DEFAULT 'PENDING' -- PENDING, APPROVED, REJECTED
);

-- ==========================================
-- FASE 2: INVENTORY SYSTEM
-- ==========================================
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sku VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    uom VARCHAR(20) NOT NULL, -- Unit of Measure
    price DECIMAL(19, 4) DEFAULT 0, -- Gunakan Decimal untuk presisi uang
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE stock_mutations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID REFERENCES products(id),
    qty DECIMAL(19, 4) NOT NULL,
    type VARCHAR(20) NOT NULL, -- 'IN' atau 'OUT'
    reference_type VARCHAR(50), -- 'ADJUSTMENT', 'PO', 'SALES'
    reference_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE stock_adjustments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID REFERENCES products(id),
    qty_diff DECIMAL(19, 4) NOT NULL, -- Selisih stok
    reason TEXT,
    status VARCHAR(20) DEFAULT 'DRAFT', -- DRAFT -> WAITING_APPROVAL -> COMPLETED
    created_by UUID NOT NULL -- User ID dari Keycloak
);

CREATE TABLE suppliers (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE,
    address TEXT,
    phone VARCHAR(20),
    email VARCHAR(100)
);

CREATE TABLE warehouses (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE
);

