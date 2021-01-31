-- +migrate Up
CREATE TABLE IF NOT EXISTS "asset_warehouses" (
    "id" UUID NOT NULL PRIMARY KEY,
    "asset_id" UUID NOT NULL,
    "warehouse_id" UUID NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "asset_warehouses";
