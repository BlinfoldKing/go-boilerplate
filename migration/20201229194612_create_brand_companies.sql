-- +migrate Up
CREATE TABLE IF NOT EXISTS "brand_companies" (
    "id" UUID NOT NULL PRIMARY KEY,
    "brand_id" UUID NOT NULL,
    "company_id" UUID NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "brand_companies";