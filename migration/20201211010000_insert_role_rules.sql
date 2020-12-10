-- +migrate Up
INSERT INTO "casbin_rule" VALUES ('p', 'admin', '/ping/authorized', 'GET');
INSERT INTO "casbin_rule" VALUES ('p', 'public', '/ping', 'GET');

-- +migrate Down
DELETE FROM "casbin_rule";
