-- +migrate Up
CREATE TABLE IF NOT EXISTS "sensor_logs" (
    "id" UUID NOT NULL PRIMARY KEY,
    "sensor_id" UUID,
    "unit" VARCHAR,
    "payload" JSONB,
    "value" VARCHAR,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "sensor_logs";
