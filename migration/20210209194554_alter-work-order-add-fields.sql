-- +migrate Up
ALTER TABLE "work_orders"
ADD COLUMN IF NOT EXISTS "no_order" VARCHAR UNIQUE,
ADD COLUMN IF NOT EXISTS "previous_site_id" UUID,
ADD COLUMN IF NOT EXISTS "mutation_requested_by" UUID,
ADD COLUMN IF NOT EXISTS "mutation_requested_at" TIMESTAMP,
ADD COLUMN IF NOT EXISTS "mutation_approved_by" UUID,
ADD COLUMN IF NOT EXISTS "mutation_approved_at" TIMESTAMP,
ADD COLUMN IF NOT EXISTS "verified_by" UUID,
ADD COLUMN IF NOT EXISTS "verified_at" TIMESTAMP;


-- +migrate Down
ALTER TABLE "work_orders"
DROP COLUMN IF EXISTS "no_order",
DROP COLUMN IF EXISTS "previous_site_id",
DROP COLUMN IF EXISTS "mutation_requested_by",
DROP COLUMN IF EXISTS "mutation_requested_at",
DROP COLUMN IF EXISTS "mutation_approved_by",
DROP COLUMN IF EXISTS "mutation_approved_at",
DROP COLUMN IF EXISTS "verified_by",
DROP COLUMN IF EXISTS "verified_at";
