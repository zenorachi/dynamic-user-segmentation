CREATE TYPE session_type AS (
    refresh_token   VARCHAR(255),
    expires_at      TIMESTAMP
);

CREATE TABLE users (
    id              SERIAL PRIMARY KEY,
    login           VARCHAR(255) UNIQUE NOT NULL,
    email           VARCHAR(255) UNIQUE NOT NULL,
    password        VARCHAR(255) NOT NULL,
    session         session_type,
    registered_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE segments (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255) UNIQUE NOT NULL
);

CREATE TYPE status_type AS ENUM ('active', 'inactive');

CREATE TABLE relations (
    user_id         INT REFERENCES users(id) ON DELETE CASCADE,
    segment_id      INT REFERENCES segments(id) ON DELETE CASCADE,
--     date_added      TIMESTAMP NOT NULL DEFAULT NOW(),
--     expires_at      TIMESTAMP DEFAULT NULL,
--     status          status_type NOT NULL DEFAULT 'active',
    PRIMARY KEY     (user_id, segment_id)
);

CREATE TYPE operation_type AS ENUM ('added', 'deleted');

CREATE TABLE operations (
    id              SERIAL PRIMARY KEY,
    user_id         INT NOT NULL,
    segment_name    VARCHAR(255) NOT NULL,
    type            operation_type NOT NULL,
    date            TIMESTAMP NOT NULL DEFAULT NOW()
)