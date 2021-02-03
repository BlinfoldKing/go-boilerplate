-- +migrate Up
ALTER TABLE "work_orders"
ADD COLUMN IF NOT EXISTS "site_id" UUID;

-- +migrate Down
ALTER TABLE "work_orders"
DROP COLUMN IF EXISTS "site_id";
