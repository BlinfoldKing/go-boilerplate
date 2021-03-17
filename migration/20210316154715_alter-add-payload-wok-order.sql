-- +migrate Up
CREATE TABLE IF NOT EXISTS "new_table" (
    "id" UUID NOT NULL PRIMARY KEY,
);

-- +migrate Down
DROP TABLE IF EXISTS "new_table";