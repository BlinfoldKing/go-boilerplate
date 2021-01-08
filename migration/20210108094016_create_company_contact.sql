-- +migrate Up
CREATE TABLE IF NOT EXISTS "company_contacts" (
    "id" UUID NOT NULL PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "contact_id" UUID NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "company_contacts";