DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE users
(
    user_id    UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    login      VARCHAR(32)              NOT NULL CHECK ( login <> '' ),
    password   VARCHAR(250)             NOT NULL CHECK ( octet_length(password) <> 0 )
);

CREATE TABLE accounts
(
    user_id    VARCHAR,
    title      VARCHAR(32)              NOT NULL,
    login      bytea                     NOT NULL,
    password       bytea                 NOT NULL
);

CREATE TABLE cards
(
    user_id         VARCHAR,
    title           VARCHAR(32)        NOT NULL,
    name           bytea                NOT NULL,
    card_number     bytea                NOT NULL,
    date_exp        bytea,
    cvc_code       bytea
);

CREATE TABLE binaries
(
    user_id    VARCHAR,
    title      VARCHAR(32)              NOT NULL  ,
    data       bytea                     NOT NULL
);

CREATE TABLE texts
(
    user_id    VARCHAR,
    title      VARCHAR(32)              NOT NULL ,
    data       bytea             NOT NULL
);
