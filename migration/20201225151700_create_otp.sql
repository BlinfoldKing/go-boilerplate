-- +migrate Up
CREATE TABLE IF NOT EXISTS "otps" (
    "token" VARCHAR NOT NULL PRIMARY KEY,
    "email" TEXT NOT NULL,
    "purpose" INT NOT NULL,
    "expired_at" TIMESTAMP WITH TIME ZONE,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "otps";
