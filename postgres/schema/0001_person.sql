-- +goose Up

CREATE TABLE person_status(
    id INT NOT NULL PRIMARY KEY,
    title VARCHAR(20) NOT NULL
);

INSERT INTO person_status(id, title) VALUES (1, 'active');
INSERT INTO person_status(id, title) VALUES (2, 'inactive');
INSERT INTO person_status(id, title) VALUES (3, 'disabled');

CREATE TABLE person_details(
    id UUID NOT NULL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE person (
    id UUID NOT NULL PRIMARY KEY,
    email_address TEXT NOT NULL,
    password TEXT NOT NULL,
    status_id INT NOT NULL REFERENCES person_status(id),
    person_details_id UUID NOT NULL REFERENCES person_details(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_username_idx ON person(email_address);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_modified_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at = NOW();
    RETURN NEW;
END;
$$ language plpgsql;
-- +goose StatementEnd

CREATE TRIGGER update_person_modified_at_column BEFORE INSERT OR UPDATE ON person FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_person_details_modified_at_column BEFORE INSERT OR UPDATE ON person_details FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_person_details_modified_at_column ON person_details;
DROP TRIGGER IF EXISTS update_person_modified_at_column ON person;

DROP FUNCTION update_modified_at_column();

DROP TABLE person;
DROP TABLE person_details;
DROP TABLE person_status;
