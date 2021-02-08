-- +migrate Up
ALTER TABLE "work_order_assets"
ADD COLUMN IF NOT EXISTS "status" INT;

-- +migrate Down
ALTER TABLE "work_order_assets"
DROP COLUMN IF EXISTS "status";
