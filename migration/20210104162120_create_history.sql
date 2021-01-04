-- +migrate Up
CREATE TABLE IF NOT EXISTS "history" (
    "id" UUID NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "asset_id" UUID NOT NULL,
    "action" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    "cost" FLOAT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "history";