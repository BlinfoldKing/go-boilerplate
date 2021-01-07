-- +migrate Up
CREATE TABLE IF NOT EXISTS "involved_users" (
    "id" UUID NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "work_order_id" UUID NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "involved_users";