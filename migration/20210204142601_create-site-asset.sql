-- +migrate Up
CREATE TABLE IF NOT EXISTS "site_assets" (
    "id" UUID PRIMARY KEY NOT NULL,
    "site_id" UUID NOT NULL,
    "asset_id" UUID NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS "site_assets";
