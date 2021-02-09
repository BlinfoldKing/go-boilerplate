-- +migrate Up
CREATE TABLE IF NOT EXISTS "templates_involved_ids" (
    "id" UUID PRIMARY KEY NOT NULL,
    "templates_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS "templates_involved_ids";
