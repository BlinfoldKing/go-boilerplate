-- +migrate Up
CREATE TABLE IF NOT EXISTS "product_categories" (
    "id" UUID NOT NULL PRIMARY KEY,
    "parent_id" UUID NOT NULL,
    "code" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "product_categories";