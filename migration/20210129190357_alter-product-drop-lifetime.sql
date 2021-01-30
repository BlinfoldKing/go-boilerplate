-- +migrate Up
ALTER TABLE "products" DROP COLUMN IF EXISTS "lifetime";

-- +migrate Down
ALTER TABLE "products" ADD COLUMN IF NOT EXISTS "lifetime" TIMESTAMP WITH TIME ZONE;
