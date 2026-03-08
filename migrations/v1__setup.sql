-- public.adjustment_reasons definition

-- Drop table

-- DROP TABLE public.adjustment_reasons;

CREATE TABLE public.adjustment_reasons (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(100) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT adjustment_reasons_name_key UNIQUE (name),
	CONSTRAINT adjustment_reasons_pkey PRIMARY KEY (id)
);


-- public.approval_workflows definition

-- Drop table

-- DROP TABLE public.approval_workflows;

CREATE TABLE public.approval_workflows (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	doc_code varchar(50) NOT NULL,
	doc_name varchar(100) NOT NULL,
	is_active bool DEFAULT true NULL,
	CONSTRAINT approval_workflows_doc_code_key UNIQUE (doc_code),
	CONSTRAINT approval_workflows_pkey PRIMARY KEY (id)
);


-- public.casbin_rule definition

-- Drop table

-- DROP TABLE public.casbin_rule;

CREATE TABLE public.casbin_rule (
	id bigserial NOT NULL,
	ptype varchar(100) NULL,
	v0 varchar(100) NULL,
	v1 varchar(100) NULL,
	v2 varchar(100) NULL,
	v3 varchar(100) NULL,
	v4 varchar(100) NULL,
	v5 varchar(100) NULL,
	CONSTRAINT casbin_rule_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_casbin_rule ON public.casbin_rule USING btree (ptype, v0, v1, v2, v3, v4, v5);


-- public.categories definition

-- Drop table

-- DROP TABLE public.categories;

CREATE TABLE public.categories (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(100) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT categories_name_key UNIQUE (name),
	CONSTRAINT categories_pkey PRIMARY KEY (id)
);


-- public.product_categories definition

-- Drop table

-- DROP TABLE public.product_categories;

CREATE TABLE public.product_categories (
	product_id uuid NOT NULL,
	category_id uuid NOT NULL,
	CONSTRAINT product_categories_pkey PRIMARY KEY (product_id, category_id)
);


-- public.product_uoms definition

-- Drop table

-- DROP TABLE public.product_uoms;

CREATE TABLE public.product_uoms (
	product_id uuid NOT NULL,
	uom_id uuid NOT NULL,
	CONSTRAINT product_uoms_pkey PRIMARY KEY (product_id, uom_id)
);


-- public.stock_movements definition

-- Drop table

-- DROP TABLE public.stock_movements;

CREATE TABLE public.stock_movements (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	movement_no varchar(50) NOT NULL,
	"type" varchar(20) NOT NULL,
	origin_warehouse_id uuid NULL,
	destination_warehouse_id uuid NULL,
	reference_no varchar(100) NULL,
	status varchar(20) DEFAULT 'DRAFT' NULL,
	transaction_date date DEFAULT CURRENT_DATE NOT NULL,
	note text NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT stock_movements_movement_no_key UNIQUE (movement_no),
	CONSTRAINT stock_movements_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_movement_date ON public.stock_movements USING btree (transaction_date);
CREATE INDEX idx_movement_type_status ON public.stock_movements USING btree (type, status);


-- public.stock_mutations definition

-- Drop table

-- DROP TABLE public.stock_mutations;

CREATE TABLE public.stock_mutations (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	product_id uuid NULL,
	qty numeric(19, 4) NOT NULL,
	"type" varchar(20) NOT NULL,
	reference_type varchar(50) NULL,
	reference_id uuid NOT NULL,
	created_at timestamptz DEFAULT now() NULL,
	CONSTRAINT stock_mutations_pkey PRIMARY KEY (id)
);


-- public.suppliers definition

-- Drop table

-- DROP TABLE public.suppliers;

CREATE TABLE public.suppliers (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(255) NOT NULL,
	contact_person varchar(100) NULL,
	email varchar(150) NULL,
	phone varchar(20) NULL,
	address text NULL,
	is_active bool DEFAULT true NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	code varchar(20) NOT NULL,
	CONSTRAINT suppliers_pkey PRIMARY KEY (id),
	CONSTRAINT suppliers_code_key UNIQUE (code)
);


-- public.uoms definition

-- Drop table

-- DROP TABLE public.uoms;

CREATE TABLE public.uoms (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(20) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT uoms_name_key UNIQUE (name),
	CONSTRAINT uoms_pkey PRIMARY KEY (id)
);


-- public.warehouses definition

-- Drop table

-- DROP TABLE public.warehouses;

CREATE TABLE public.warehouses (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(100) NOT NULL,
	"location" text NULL,
	is_active bool DEFAULT true NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	code varchar(20) NOT NULL,
	CONSTRAINT warehouses_pkey PRIMARY KEY (id),
	CONSTRAINT warehouses_unique UNIQUE (code)
);


-- public.approval_steps definition

-- Drop table

-- DROP TABLE public.approval_steps;

CREATE TABLE public.approval_steps (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	workflow_id uuid NULL,
	step_order int4 NOT NULL,
	target_group_name varchar(100) NOT NULL,
	min_approver int4 DEFAULT 1 NULL,
	CONSTRAINT approval_steps_pkey PRIMARY KEY (id),
	CONSTRAINT approval_steps_workflow_id_fkey FOREIGN KEY (workflow_id) REFERENCES public.approval_workflows(id) ON DELETE CASCADE
);


-- public.approval_transactions definition

-- Drop table

-- DROP TABLE public.approval_transactions;

CREATE TABLE public.approval_transactions (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	workflow_id uuid NULL,
	reference_id uuid NOT NULL,
	current_step int4 DEFAULT 1 NULL,
	status varchar(20) DEFAULT 'PENDING'::character varying NULL,
	CONSTRAINT approval_transactions_pkey PRIMARY KEY (id),
	CONSTRAINT approval_transactions_workflow_id_fkey FOREIGN KEY (workflow_id) REFERENCES public.approval_workflows(id)
);


-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	sku varchar(50) NOT NULL,
	"name" varchar(255) NOT NULL,
	category_id uuid NULL,
	uom_id uuid NULL,
	supplier_id uuid NULL,
	price numeric(19, 4) DEFAULT 0 NULL,
	min_stock numeric(19, 4) DEFAULT 0 NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	is_active bool DEFAULT true NOT NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id),
	CONSTRAINT products_sku_key UNIQUE (sku),
	CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL,
	CONSTRAINT products_primary_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) ON DELETE SET NULL,
	CONSTRAINT products_uom_id_fkey FOREIGN KEY (uom_id) REFERENCES public.uoms(id) ON DELETE RESTRICT
);


-- public.stock_adjustments definition

-- Drop table

-- DROP TABLE public.stock_adjustments;

CREATE TABLE public.stock_adjustments (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	adjustment_no varchar(50) NOT NULL,
	warehouse_id uuid NOT NULL,
	transaction_date date DEFAULT CURRENT_DATE NOT NULL,
	status varchar(20) DEFAULT 'DRAFT'::character varying NULL,
	note text NULL,
	created_by uuid NULL,
	approved_by uuid NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT stock_adjustments_adjustment_no_key UNIQUE (adjustment_no),
	CONSTRAINT stock_adjustments_pkey PRIMARY KEY (id),
	CONSTRAINT stock_adjustments_warehouse_id_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id)
);


-- public.stock_levels definition

-- Drop table

-- DROP TABLE public.stock_levels;

CREATE TABLE public.stock_levels (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	product_id uuid NOT NULL,
	warehouse_id uuid NOT NULL,
	quantity numeric(19, 4) DEFAULT 0 NOT NULL,
	last_updated timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT stock_levels_pkey PRIMARY KEY (id),
	CONSTRAINT stock_levels_product_id_warehouse_id_key UNIQUE (product_id, warehouse_id),
	CONSTRAINT stock_levels_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE,
	CONSTRAINT stock_levels_warehouse_id_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id) ON DELETE CASCADE
);
CREATE INDEX idx_stock_levels_product ON public.stock_levels USING btree (product_id);


-- public.stock_movement_items definition

-- Drop table

-- DROP TABLE public.stock_movement_items;

CREATE TABLE public.stock_movement_items (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	stock_movement_id uuid NOT NULL,
	product_id uuid NOT NULL,
	quantity numeric(19, 4) NOT NULL,
	note text NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT stock_movement_items_pkey PRIMARY KEY (id),
	CONSTRAINT stock_movement_items_movement_id_fkey FOREIGN KEY (stock_movement_id) REFERENCES public.stock_movements(id) ON DELETE CASCADE
);


-- public.stock_transactions definition

-- Drop table

-- DROP TABLE public.stock_transactions;

CREATE TABLE public.stock_transactions (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	product_id uuid NOT NULL,
	warehouse_id uuid NOT NULL,
	supplier_id uuid NULL,
	"type" varchar(20) NOT NULL,
	quantity numeric(19, 4) NOT NULL,
	balance_after numeric(19, 4) DEFAULT 0 NOT NULL,
	reference_no varchar(100) NULL,
	note text NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT stock_transactions_pkey PRIMARY KEY (id),
	CONSTRAINT stock_transactions_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id),
	CONSTRAINT stock_transactions_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id),
	CONSTRAINT stock_transactions_warehouse_id_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id)
);
CREATE INDEX idx_stock_transactions_product_date ON public.stock_transactions USING btree (product_id, created_at);

-- Table Triggers

create trigger trg_before_stock_transaction before
insert
    on
    public.stock_transactions for each row execute function fn_before_insert_transaction();


-- public.stock_adjustment_items definition

-- Drop table

-- DROP TABLE public.stock_adjustment_items;

CREATE TABLE public.stock_adjustment_items (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	stock_adjustment_id uuid NOT NULL,
	product_id uuid NOT NULL,
	reason_id uuid NULL,
	system_qty numeric(19, 4) NOT NULL,
	actual_qty numeric(19, 4) NOT NULL,
	adjustment_qty numeric(19, 4) NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	CONSTRAINT stock_adjustment_items_pkey PRIMARY KEY (id),
	CONSTRAINT stock_adjustment_items_adjustment_id_fkey FOREIGN KEY (stock_adjustment_id) REFERENCES public.stock_adjustments(id) ON DELETE CASCADE,
	CONSTRAINT stock_adjustment_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id),
	CONSTRAINT stock_adjustment_items_reason_id_fkey FOREIGN KEY (reason_id) REFERENCES public.adjustment_reasons(id)
);