-- USER'S SESSION --
CREATE TYPE session_type AS (
    refresh_token   VARCHAR(255),
    expires_at      TIMESTAMP
);

-- USERS --
CREATE TABLE users (
    id              SERIAL PRIMARY KEY,
    login           VARCHAR(255) UNIQUE NOT NULL,
    email           VARCHAR(255) UNIQUE NOT NULL,
    password        VARCHAR(255) NOT NULL,
    session         session_type,
    registered_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

-- SEGMENTS --
CREATE TABLE segments (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255) UNIQUE NOT NULL,
    assign_percent  NUMERIC(10, 2) DEFAULT 0.0
);

-- USER-SEGMENT RELATIONS (MANY-TO-MANY) --
-- DELETE REFERENCES TO INSERT INTO OPERATIONS IF SEGMENT WAS DELETED --
CREATE TABLE relations (
    user_id         INT NOT NULL,
    segment_id      INT NOT NULL,
    PRIMARY KEY     (user_id, segment_id)
);

-- OPERATION TYPE --
CREATE TYPE operation_type AS ENUM (
    'added', 'deleted'
);

-- OPERATIONS --
CREATE TABLE operations (
    id              SERIAL PRIMARY KEY,
    user_id         INT NOT NULL,
    segment_name    VARCHAR(255) NOT NULL,
    type            operation_type NOT NULL,
    date            TIMESTAMP NOT NULL DEFAULT NOW()
);

-- AUTOMATIC DELETION RECORDS IN relations & AUTOMATIC INSERTION INTO operations --
CREATE OR REPLACE FUNCTION segment_deleted_trigger()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD IS NOT NULL THEN
        IF EXISTS (SELECT FROM relations WHERE segment_id = OLD.id) THEN
            INSERT INTO operations (user_id, segment_name, type)
            SELECT user_id, OLD.name, 'deleted' FROM relations WHERE segment_id = OLD.id;
    END IF;

    DELETE FROM relations WHERE segment_id = OLD.id;
END IF;
RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- TRIGGER --
CREATE TRIGGER segment_deleted
    AFTER DELETE ON segments
    FOR EACH ROW
    EXECUTE FUNCTION segment_deleted_trigger();

-- AUTOMATIC ASSIGN USERS BY AUTO-ASSIGN-PERCENT --
CREATE OR REPLACE FUNCTION auto_assign_users_trigger()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.assign_percent > 0.0 THEN
        DECLARE
            total_users INT;
            users_to_assign INT;
            usr_id INT;
        BEGIN
            SELECT COUNT(*) INTO total_users FROM users;

            users_to_assign := CEIL(NEW.assign_percent * total_users);

            FOR usr_id IN SELECT id FROM users ORDER BY random() LIMIT users_to_assign
            LOOP
                INSERT INTO relations (user_id, segment_id) VALUES (usr_id, NEW.id);

                INSERT INTO operations (user_id, segment_name, type) VALUES (usr_id, NEW.name, 'added');
            END LOOP;
        END;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- TRIGGER --
CREATE TRIGGER auto_assign_users
    AFTER INSERT ON segments
    FOR EACH ROW
    EXECUTE FUNCTION auto_assign_users_trigger();
