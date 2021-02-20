-- +migrate Up
ALTER TABLE "products"
ADD COLUMN IF NOT EXISTS "salvage_value" FLOAT;

-- +migrate Down
ALTER TABLE "products"
DROP COLUMN IF EXISTS "salvage_value";
