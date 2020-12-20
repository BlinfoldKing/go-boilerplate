-- +migrate Up
CREATE TABLE IF NOT EXISTS "roles" (
    "id" UUID NOT NULL PRIMARY KEY,
    "slug" VARCHAR(255) UNIQUE NOT NULL,
    "description" TEXT NOT NULL DEFAULT 'new role'
);

-- +migrate Down
DROP TABLE IF EXISTS "roles";
