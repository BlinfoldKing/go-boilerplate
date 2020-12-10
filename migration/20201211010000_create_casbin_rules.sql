-- +migrate Up
CREATE TABLE IF NOT EXISTS "casbin_rule" (
    "p_type" TEXT NOT NULL,
    "sub" TEXT NOT NULL,
    "obj" TEXT NOT NULL,
    "act" TEXT NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "casbin_rule";
