-- +migrate Up
CREATE TABLE "public"."phones" (
    id serial NOT NULL,
    mac character varying NOT NULL UNIQUE,
    vendor character varying NOT NULL,
    model character varying NOT NULL,
    description character varying NOT NULL,
    last_seen timestamp without time zone,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
)
;

-- +migrate Down
DROP TABLE "public"."phones";