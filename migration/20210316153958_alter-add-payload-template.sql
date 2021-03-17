-- +migrate Up
ALTER TABLE "templates" ADD COLUMN IF NOT EXISTS "payload" JSONB;

-- +migrate Down
ALTER TABLE "templates" DROP COLUMN IF EXISTS "payload";

