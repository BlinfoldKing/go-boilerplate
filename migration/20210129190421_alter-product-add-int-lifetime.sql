-- +migrate Up
ALTER TABLE "products" ADD COLUMN IF NOT EXISTS "lifetime" INT;

-- +migrate Down
ALTER TABLE "products" DROP COLUMN IF EXISTS "lifetime";
