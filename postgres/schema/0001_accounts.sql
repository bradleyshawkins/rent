-- +goose Up

CREATE TABLE role (
    id INT NOT NULL PRIMARY KEY,
    role TEXT NOT NULL
);

INSERT INTO role VALUES (1, 'admin');
INSERT INTO role VALUES (2, 'write');
INSERT INTO role VALUES (3, 'read');

CREATE TABLE person (
    id UUID NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email_address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_username_idx ON person(username);
CREATE UNIQUE INDEX user_email_address_idx ON person(email_address);

CREATE TABLE account (
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE membership (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES person(id) ON UPDATE CASCADE,
    account_id UUID NOT NULL REFERENCES account(id) ON UPDATE CASCADE,
    role_id INT NOT NULL REFERENCES role(id)
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
CREATE TRIGGER update_account_modified_at_column BEFORE INSERT OR UPDATE ON account FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_membership_modified_at_column BEFORE INSERT OR UPDATE ON membership FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER update_person_modified_at_column;
DROP TRIGGER update_account_modified_at_column;
DROP TRIGGER update_membership_modified_at_column;

DROP FUNCTION update_modified_at_column();

DROP TABLE membership;
DROP TABLE account;
DROP TABLE person;
DROP TABLE role;