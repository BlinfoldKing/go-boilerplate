-- +migrate Up
ALTER TABLE "assets"
ADD COLUMN IF NOT EXISTS "salvage_value" FLOAT;

-- +migrate Down
ALTER TABLE "assets"
DROP COLUMN IF EXISTS "salvage_value";
