-- +migrate Up
CREATE TABLE IF NOT EXISTS "brands" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "origin_country" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "brands";