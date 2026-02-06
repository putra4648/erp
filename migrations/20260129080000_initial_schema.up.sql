-- EKSTENSI UNTUK UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ==========================================
-- FASE 1: APPROVAL SYSTEM
-- ==========================================
CREATE TABLE IF NOT EXISTS approval_workflows (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    doc_code VARCHAR(50) UNIQUE NOT NULL, -- Contoh: 'STOCK_ADJ'
    doc_name VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS approval_steps (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workflow_id UUID REFERENCES approval_workflows(id),
    step_order INT NOT NULL,
    target_group_name VARCHAR(100) NOT NULL, -- Nama Group dari Keycloak
    min_approver INT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS approval_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workflow_id UUID REFERENCES approval_workflows(id),
    reference_id UUID NOT NULL, -- ID dari tabel Bisnis (misal: stock_adjustment_id)
    current_step INT DEFAULT 1,
    status VARCHAR(20) DEFAULT 'PENDING' -- PENDING, APPROVED, REJECTED
);

-- ==========================================
-- FASE 2: INVENTORY & PRODUCT SYSTEM
-- ==========================================

-- Master Tables
CREATE TABLE IF NOT EXISTS uoms (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Drop existing products table to recreate
DROP TABLE IF EXISTS products CASCADE;

-- Products Table
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    sku VARCHAR(100) UNIQUE NOT NULL,
    price DECIMAL(19, 2) NOT NULL,
    cost DECIMAL(19, 2) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- Conjunction Table for Products and Categories
CREATE TABLE IF NOT EXISTS product_categories (
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    category_id UUID REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);

-- Conjunction Table for Products and UOMs
CREATE TABLE IF NOT EXISTS product_uoms (
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    uom_id UUID REFERENCES uoms(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, uom_id)
);


CREATE TABLE IF NOT EXISTS stock_mutations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID REFERENCES products(id),
    qty DECIMAL(19, 4) NOT NULL,
    type VARCHAR(20) NOT NULL, -- 'IN' atau 'OUT'
    reference_type VARCHAR(50), -- 'ADJUSTMENT', 'PO', 'SALES'
    reference_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS stock_adjustments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID REFERENCES products(id),
    qty_diff DECIMAL(19, 4) NOT NULL, -- Selisih stok
    reason TEXT,
    status VARCHAR(20) DEFAULT 'DRAFT', -- DRAFT -> WAITING_APPROVAL -> COMPLETED
    created_by UUID NOT NULL -- User ID dari Keycloak
);

CREATE TABLE IF NOT EXISTS suppliers (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE,
    address TEXT,
    phone VARCHAR(20),
    email VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS warehouses (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE
);
