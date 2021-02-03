-- +migrate Up
ALTER TABLE "users"
ADD COLUMN IF NOT EXISTS "company_contact_id" UUID;

-- +migrate Down
ALTER TABLE "users"
DROP COLUMN IF EXISTS "company_contact_id";
