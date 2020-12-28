-- +migrate Up
CREATE TABLE IF NOT EXISTS "notifications" (
    "id" UUID NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "title" TEXT,
    "subtitle" TEXT,
    "url_link" TEXT,
    "body" TEXT,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "notifications";