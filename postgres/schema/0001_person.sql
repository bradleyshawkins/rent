-- +goose Up

CREATE TABLE person (
    id UUID NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    is_active bool NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_username_idx ON person(username);

CREATE TABLE person_details(
    id UUID NOT NULL PRIMARY KEY,
    person_id UUID NOT NULL REFERENCES person,
    address_id UUID NOT NULL REFERENCES address,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email_address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_email_address_idx ON person_details(email_address);

CREATE TABLE address(
    id UUID NOT NULL PRIMARY KEY,
    street_1 TEXT NOT NULL,
    street_2 TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
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
CREATE TRIGGER update_address_modified_at_column BEFORE INSERT OR UPDATE ON address FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_person_modified_at_column ON person;
DROP TRIGGER IF EXISTS update_person_details_modified_at_column ON person_details;
DROP TRIGGER IF EXISTS update_address_modified_at_column ON address;

DROP FUNCTION update_modified_at_column();

DROP TABLE address;
DROP TABLE person_details;
DROP TABLE person;
