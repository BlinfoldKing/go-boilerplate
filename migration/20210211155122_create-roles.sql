-- +migrate Up
INSERT INTO "roles"(id, slug, description) VALUES
    ('0ee7751e-4aac-4134-b8ed-16dc57a99f58', 'pmu', 'default pmu'),
    ('961bd6bb-b588-4e63-86b0-76f81d5bb6b2', 'penyedia', 'default penyedia'),
    ('1a2c267f-16d0-4b72-878d-914c2fa3bfd6', 'auditor', 'default auditor'),
    ('f25d7a24-11b5-4922-95fb-7d2cb5ffd4f3', 'assessor', 'default assessor'),
    ('9fb1295f-b78c-445c-bf8f-7e3141d1e349', 'field-engineer', 'field-engineer');

-- +migrate Down
DELETE FROM "roles" WHERE
slug = 'pmu'
OR slug = 'penyedia'
OR slug = 'auditor'
OR slug = 'assessor'
OR slug = 'field-engineer';
