-- +migrate Up
ALTER TABLE "work_order_assets"
ADD COLUMN IF NOT EXISTS "edited_by" UUID;

-- +migrate Down
ALTER TABLE "work_order_assets"
DROP COLUMN IF EXISTS "edited_by";
