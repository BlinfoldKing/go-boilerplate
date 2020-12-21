-- +migrate Up
INSERT INTO "roles"(id, slug, description) VALUES
    ('127805c7-9d1a-4332-8786-3b988da607e2', 'admin', 'default admin'),
    ('9d02fc35-3d92-4755-bbe2-8ba99d2b57b2', 'member', 'default member');

-- +migrate Down
DELETE FROM "roles" WHERE slug = 'admin' OR slug = 'member';
