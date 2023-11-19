CREATE SCHEMA IF NOT EXISTS user_schema;

CREATE TABLE user_schema.user
(
    id             uuid,
    username       text,
    password       text,
    lastLogin      date,
    PRIMARY KEY (id)
);

CREATE TABLE user_schema.role (
    id uuid,
    name text,
    description text,
    PRIMARY KEY (id)
);

CREATE TABLE user_schema.user_role (
    id uuid,
    roleId uuid,
    userId uuid,
    PRIMARY KEY (id),
    CONSTRAINT fk_user_schema_user
               FOREIGN KEY (userId)
               REFERENCES user_schema.user (id),
    CONSTRAINT fk_role_schema_user
        FOREIGN KEY (roleId)
            REFERENCES user_schema.role (id)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";