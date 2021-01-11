-- +migrate Up
ALTER TABLE "otps"
ALTER COLUMN token TYPE VARCHAR;
