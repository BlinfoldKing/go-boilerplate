-- +migrate Up
ALTER TABLE "work_orders"
ADD COLUMN IF NOT EXISTS "status" INT;

-- +migrate Down
ALTER TABLE "work_orders"
DROP COLUMN IF EXISTS "status";
