-- +migrate Up
CREATE TABLE IF NOT EXISTS "work_order_products" (
    "id" UUID NOT NULL PRIMARY KEY,
    "work_order_id" UUID NOT NULL,
    "product_id" UUID NOT NULL,
    "qty" INT NOT NULL,
    "status" INT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "work_order_products";
