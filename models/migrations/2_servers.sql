-- +migrate Up
CREATE TABLE "public"."servers" (
    id serial NOT NULL,
    url character varying NOT NULL UNIQUE,
    description character varying NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
)
;

-- +migrate Down
DROP TABLE "public"."servers";