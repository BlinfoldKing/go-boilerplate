-- +migrate Up
CREATE TABLE IF NOT EXISTS "work_orders" (
    "id" UUID NOT NULL PRIMARY KEY,
    "pic_id" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "type" INT NOT NULL,
    "description" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "work_orders";