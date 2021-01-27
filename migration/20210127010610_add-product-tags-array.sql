-- +migrate Up
ALTER TABLE "products" ADD COLUMN IF NOT EXISTS "tags" VARCHAR[];

-- +migrate Down
ALTER TABLE "products" DROP COLUMN IF EXISTS "tags";

