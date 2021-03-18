-- +migrate Up
CREATE TABLE IF NOT EXISTS "sensors" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR,
    "description" TEXT,
    "site_id" UUID,
    "sensor_type" INT,
    "code" VARCHAR,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "sensors";
