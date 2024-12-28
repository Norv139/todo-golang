CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE todo (
    "createdAt"     TIMESTAMP NOT NULL DEFAULT now(),
    "updatedAt"     TIMESTAMP NOT NULL DEFAULT now(),
    "deletedAt"     TIMESTAMP,
    "uid"           uuid NOT NULL DEFAULT uuid_generate_v4(),
    "desc"          character varying(255),
    "name"          character varying(255),
    "check"         BOOLEAN DEFAULT FALSE
);