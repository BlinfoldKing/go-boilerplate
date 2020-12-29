-- +migrate Up
ALTER TABLE "users" ADD COLUMN IF NOT EXISTS "active_status" INT DEFAULT 0;

-- +migrate Down
ALTER TABLE "users" DROP COLUMN IF EXISTS "active_status";