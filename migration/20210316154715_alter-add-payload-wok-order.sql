-- +migrate Up
ALTER TABLE "work_orders" ADD COLUMN IF NOT EXISTS "payload" JSONB;

-- +migrate Down
ALTER TABLE "work_orders" DROP COLUMN IF EXISTS "payload";

