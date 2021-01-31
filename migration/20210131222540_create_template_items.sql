-- +migrate Up
CREATE TABLE IF NOT EXISTS "template_items" (
    "id" UUID NOT NULL PRIMARY KEY,
    "template_id" UUID NOT NULL,
    "product_id" UUID NOT NULL,
    "qty" INT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "template_items";