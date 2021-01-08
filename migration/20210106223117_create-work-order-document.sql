-- +migrate Up
CREATE TABLE IF NOT EXISTS "work_order_documents" (
    "id" UUID NOT NULL PRIMARY KEY,
    "work_order_id" UUID NOT NULL,
    "document_id" UUID NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "deleted_at" TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS "work_order_documents";