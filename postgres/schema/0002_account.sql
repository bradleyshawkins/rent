-- +goose Up

CREATE TABLE account_status(
    id INT NOT NULL PRIMARY KEY,
    value TEXT NOT NULL
);

INSERT INTO account_status(id, value) VALUES (1, 'active');
INSERT INTO account_status(id, value) VALUES (2, 'disabled');
INSERT INTO account_status(id, value) VALUES (3, 'canceled');

CREATE TABLE account(
    id UUID NOT NULL PRIMARY KEY,
    status INT NOT NULL REFERENCES account_status(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_account_modified_at_column BEFORE INSERT OR UPDATE ON account FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

CREATE TABLE role (
    id INT NOT NULL PRIMARY KEY,
    value TEXT NOT NULL
);

INSERT INTO role(id, value) VALUES (1, 'owner');
INSERT INTO role(id, value) VALUES (2, 'write');
INSERT INTO role(id, value) VALUES (3, 'read');

CREATE TABLE membership (
    account_id UUID NOT NULL REFERENCES account(id),
    person_id UUID NOT NULL REFERENCES person(id),
    role_id INT NOT NULL REFERENCES role(id),
    PRIMARY KEY (account_id, person_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_membership_modified_at_column BEFORE INSERT OR UPDATE ON membership FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER update_membership_modified_at_column ON membership;
DROP TRIGGER update_account_modified_at_column ON account;

DROP TABLE membership;
DROP TABLE role;
DROP TABLE account;
DROP TABLE account_status;