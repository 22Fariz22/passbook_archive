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
    login       VARCHAR(32)             NOT NULL,
    password       VARCHAR(250)         NOT NULL
);

CREATE TABLE cards
(
    user_id         VARCHAR,
    title           VARCHAR(32)        NOT NULL,
    card_number     VARCHAR(32)        NOT NULL,
    name            VARCHAR(32),
    date_exp        VARCHAR(6),
    cvc_code        VARCHAR(6)
);

CREATE TABLE binaries
(
    user_id    VARCHAR,
    title      VARCHAR(32)              NOT NULL  ,
    data       VARCHAR(10000)           NOT NULL
);

CREATE TABLE texts
(
    user_id    VARCHAR,
    title      VARCHAR(32)              NOT NULL ,
    data       VARCHAR(1000)             NOT NULL
);
