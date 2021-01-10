-- +migrate Up
ALTER TABLE phones ADD COLUMN fanvil_tz integer;

-- +migrate Down
ALTER TABLE phones DROP COLUMN IF EXISTS fanvil_tz;