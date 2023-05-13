DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
--CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


--CREATE TYPE role AS ENUM ('admin', 'user');

CREATE TABLE users
(
    user_id    UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    login VARCHAR(32)              NOT NULL CHECK ( first_name <> '' ),
    password   VARCHAR(250)             NOT NULL CHECK ( octet_length(password) <> 0 ),
);