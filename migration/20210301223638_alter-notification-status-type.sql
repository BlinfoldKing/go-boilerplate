-- +migrate Up
ALTER TABLE "notifications"
ADD COLUMN IF NOT EXISTS "status" INT,
ADD COLUMN IF NOT EXISTS "type" INT;


-- +migrate Down
ALTER TABLE "notifications"
DROP COLUMN IF EXISTS "status",
DROP COLUMN IF EXISTS "type";
