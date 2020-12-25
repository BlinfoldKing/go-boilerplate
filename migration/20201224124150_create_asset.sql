-- +migrate Up
CREATE TABLE IF NOT EXISTS "assets" (
    "id" UUID NOT NULL PRIMARY KEY,
    "product_id" uuid NOT NULL,
    "serial_number" VARCHAR,
    "status" INT,
    "purchase_date" TIMESTAMP,
    "purchase_price" FLOAT,
    "supplier_company_id" UUID NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    "order" SERIAL NOT NULL
);

ALTER TABLE "assets" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "assets" ADD FOREIGN KEY ("supplier_company_id") REFERENCES "companies" ("id");

-- +migrate Down
DROP TABLE IF EXISTS "assets";
