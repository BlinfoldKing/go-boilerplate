-- +migrate Up
CREATE TABLE IF NOT EXISTS "product_specifications" (
    "id" UUID NOT NULL PRIMARY KEY,
    "product_id" UUID NOT NULL,
    "parameter" TEXT NOT NULL,
    "value" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "product_specifications";