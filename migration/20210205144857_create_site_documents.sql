-- +migrate Up
CREATE TABLE IF NOT EXISTS "site_documents" (
    "id" UUID PRIMARY KEY NOT NULL,
    "site_id" UUID NOT NULL,
    "document_id" UUID NOT NULL,
    "approve_status" INT,
    "notes" VARCHAR,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP

);

-- +migrate Down
DROP TABLE IF EXISTS "site_documents";
