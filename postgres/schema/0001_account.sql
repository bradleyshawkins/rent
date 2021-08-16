-- +goose Up

CREATE TABLE person_status(
    id INT NOT NULL PRIMARY KEY,
    title VARCHAR(20) NOT NULL
);

INSERT INTO person_status(id, title) VALUES (1, 'active');
INSERT INTO person_status(id, title) VALUES (2, 'inactive');
INSERT INTO person_status(id, title) VALUES (3, 'disabled');

CREATE TABLE person (
    id UUID NOT NULL PRIMARY KEY,
    email_address TEXT NOT NULL,
    password TEXT NOT NULL,
    status_id INT NOT NULL REFERENCES person_status(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_username_idx ON person(email_address);

CREATE TABLE person_details(
    id UUID NOT NULL PRIMARY KEY,
    person_id UUID NOT NULL REFERENCES person(id),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_status(
    id INT NOT NULL PRIMARY KEY,
    title VARCHAR(20) NOT NULL
);

INSERT INTO account_status(id, title) VALUES(1, 'active');
INSERT INTO account_status(id, title) VALUES(2, 'inactive');
INSERT INTO account_status(id, title) VALUES(3, 'disabled');

CREATE TABLE IF NOT EXISTS account(
    id UUID NOT NULL PRIMARY KEY,
    status_id INT NOT NULL REFERENCES account_status(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS membership(
    person_id UUID NOT NULL REFERENCES person(id),
    account_id UUID NOT NULL REFERENCES account(id),
    permission INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (person_id, account_id)
);

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
CREATE TRIGGER update_account_modified_at_column BEFORE INSERT OR UPDATE ON account FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_membership_modified_at_column BEFORE INSERT OR UPDATE ON membership FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_membership_modified_at_column ON membership;
DROP TRIGGER IF EXISTS update_account_modified_at_column ON account;
DROP TRIGGER IF EXISTS update_person_details_modified_at_column ON person_details;
DROP TRIGGER IF EXISTS update_person_modified_at_column ON person;

DROP FUNCTION update_modified_at_column();

DROP TABLE account;
DROP TABLE person_details;
DROP TABLE person;
