-- +migrate Up
CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID NOT NULL PRIMARY KEY,
    "email" TEXT UNIQUE NOT NULL,
    "password_hash" TEXT NOT NULL
);


-- +migrate Down
DROP TABLE IF EXISTS "users";
