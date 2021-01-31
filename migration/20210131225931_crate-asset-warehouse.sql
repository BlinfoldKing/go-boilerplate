-- +migrate Up
CREATE TABLE IF NOT EXISTS "asset_warehouses" (
    "id" UUID NOT NULL PRIMARY KEY,
    "asset_id" UUID NOT NULL PRIMARY KEY,
    "warehouse_id" UUID NOT NULL PRIMARY KEY,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
);

-- +migrate Down
DROP TABLE IF EXISTS "asset_warehouses";
