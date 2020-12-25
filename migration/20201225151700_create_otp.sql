-- +migrate Up
CREATE TABLE IF NOT EXISTS "otps" (
    "token" UUID NOT NULL PRIMARY KEY,
    "email" TEXT NOT NULL,
    "purpose" int NOT NULL,
    "expired_at" TIMESTAMP WITH TIME ZONE,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "otps";
