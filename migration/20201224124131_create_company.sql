-- +migrate Up
CREATE TABLE IF NOT EXISTS "companies" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "companies";
