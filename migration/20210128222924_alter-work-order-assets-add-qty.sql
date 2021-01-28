-- +migrate Up
ALTER TABLE "work_order_assets"
ADD COLUMN IF NOT EXISTS "qty" INT;

-- +migrate Down
ALTER TABLE "work_order_assets"
DROP COLUMN IF EXISTS "qty";
