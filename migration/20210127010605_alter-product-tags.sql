-- +migrate Up
ALTER TABLE "products" DROP COLUMN IF EXISTS "tags";

-- +migrate Down
ALTER TABLE "products" ADD COLUMN IF NOT EXISTS "tags" VARCHAR;

