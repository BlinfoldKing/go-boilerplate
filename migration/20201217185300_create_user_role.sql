-- +migrate Up
CREATE TABLE IF NOT EXISTS "user_roles" (
    "id" UUID NOT NULL PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "role_id" UUID NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "user_roles";
