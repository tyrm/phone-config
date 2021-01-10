-- +migrate Up
CREATE TABLE "public"."extensions" (
    id serial NOT NULL,
    extension integer NOT NULL,
    secret character varying NOT NULL,
    server_id integer NOT NULL,
    display_name character varying NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    UNIQUE ("extension", "server_id"),
    CONSTRAINT fk_server_id FOREIGN KEY (server_id) REFERENCES servers (id) ON DELETE RESTRICT
)
;

-- +migrate Down
DROP TABLE "public"."extensions";