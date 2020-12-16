-- +migrate Up
CREATE TABLE IF NOT EXISTS "documents" (
    "id" UUID NOT NULL PRIMARY KEY,
    "object_name" TEXT NOT NULL,
    "bucket_name" TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "documents";