-- +migrate Up
ALTER TABLE "users"
ADD COLUMN IF NOT EXISTS "created_at" TIMESTAMP WITH TIME ZONE,
ADD COLUMN IF NOT EXISTS "updated_at" TIMESTAMP WITH TIME ZONE,
ADD COLUMN IF NOT EXISTS "deleted_at" TIMESTAMP WITH TIME ZONE;

-- +migrate Down
ALTER TABLE "users"
DROP COLUMN IF EXISTS "created_at",
DROP COLUMN IF EXISTS "updated_at",
DROP COLUMN IF EXISTS "deleted_at";
