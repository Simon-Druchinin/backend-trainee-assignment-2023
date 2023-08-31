CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS segments
(
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users_segments
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    segment_id INT REFERENCES segments (id) ON DELETE CASCADE NOT NULL,
    UNIQUE(user_id, segment_id)
);

CREATE TABLE IF NOT EXISTS users_segments_history
(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    segment_id INT REFERENCES segments (id) ON DELETE CASCADE NOT NULL,
	operation_type VARCHAR(255) NOT NULL,
	timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_users_segments_history() RETURNS TRIGGER AS $users_segments_history$
    BEGIN
        IF (TG_OP = 'INSERT') THEN
            INSERT INTO users_segments_history (user_id, segment_id, operation_type) SELECT NEW.user_id, NEW.segment_id, 'INSERT';
            RETURN NEW;
		ELSIF (TG_OP = 'DELETE') THEN
            INSERT INTO users_segments_history (user_id, segment_id, operation_type) SELECT OLD.user_id, OLD.segment_id, 'DELETE';
            RETURN OLD;
		END IF;
        RETURN NULL; -- result is ignored since this is an AFTER trigger
    END;
$users_segments_history$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_users_segments_history_on_insert_trigger
AFTER INSERT OR UPDATE OR DELETE ON users_segments
    FOR EACH ROW EXECUTE PROCEDURE update_users_segments_history();

