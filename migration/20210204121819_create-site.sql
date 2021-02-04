-- +migrate Up
CREATE TABLE IF NOT EXISTS "sites" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "latitude" FLOAT NOT NULL,
    "longitude" FLOAT NOT NULL,
    "description" VARCHAR NOT NULL,
    "address" VARCHAR NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS "sites";
