-- +migrate Up
CREATE TABLE "public"."phone_lines" (
    id serial NOT NULL,
    phone_id integer NOT NULL,
    extension_id integer NOT NULL,
    position integer NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    UNIQUE ("phone_id", "position"),
    CONSTRAINT fk_phone_id FOREIGN KEY (phone_id) REFERENCES phones (id) ON DELETE RESTRICT,
    CONSTRAINT fk_server_id FOREIGN KEY (extension_id) REFERENCES extensions (id) ON DELETE RESTRICT
)
;

-- +migrate Down
DROP TABLE "public"."phone_lines";