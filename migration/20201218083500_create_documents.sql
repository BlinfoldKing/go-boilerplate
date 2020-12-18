-- +migrate Up
CREATE TABLE IF NOT EXISTS "documents" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "file_size" INT NOT NULL,
    "file_type" TEXT NOT NULL,
    "object_name" TEXT NOT NULL,
    "bucket_name" TEXT NOT NULL,
    "url_link" TEXT,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "documents";