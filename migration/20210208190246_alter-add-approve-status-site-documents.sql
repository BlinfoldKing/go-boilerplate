-- +migrate Up
ALTER TABLE "site_documents"
ADD COLUMN IF NOT EXISTS "approve_status" INT,
ADD COLUMN IF NOT EXISTS "notes" VARCHAR;

-- +migrate Down
ALTER TABLE "site_documents"
DROP COLUMN IF EXISTS "approve_status",
DROP COLUMN IF EXISTS "notes";
