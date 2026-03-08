-- DROP FUNCTION public.fn_before_insert_transaction();

CREATE OR REPLACE FUNCTION public.fn_before_insert_transaction()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
DECLARE
    latest_stock DECIMAL(19, 4);
BEGIN
    -- Ambil quantity terbaru dari tabel stock_levels
    SELECT quantity INTO latest_stock
    FROM stock_levels
    WHERE product_id = NEW.product_id 
      AND warehouse_id = NEW.warehouse_id;

    -- Isi nilai NEW.balance_after sebelum insert terjadi
    NEW.balance_after := COALESCE(latest_stock, 0);

    RETURN NEW;
END;
$function$
;

-- DROP FUNCTION public.fn_update_balance_after();

CREATE OR REPLACE FUNCTION public.fn_update_balance_after()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
DECLARE
    current_stock DECIMAL(19, 4);
BEGIN
    -- Ambil saldo terbaru dari tabel stock_levels untuk produk dan gudang terkait
    SELECT quantity INTO current_stock
    FROM stock_levels
    WHERE product_id = NEW.product_id 
      AND warehouse_id = NEW.warehouse_id;

    -- Update kolom balance_after pada baris transaksi yang baru saja dimasukkan
    UPDATE stock_transactions
    SET balance_after = COALESCE(current_stock, 0)
    WHERE id = NEW.id;

    RETURN NEW;
END;
$function$
;
