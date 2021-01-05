-- +migrate Up
ALTER TABLE "products"
ADD COLUMN IF NOT EXISTS "type" VARCHAR;

-- +migrate Down
ALTER TABLE "products"
DROP COLUMN IF EXISTS "type";