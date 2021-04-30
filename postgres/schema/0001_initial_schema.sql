-- +goose Up
CREATE TABLE IF NOT EXISTS address(
    id UUID NOT NULL PRIMARY KEY,
    address_1 TEXT NOT NULL,
    address_2 TEXT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zipcode VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE person (
    id UUID NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL, -- todo add better password stuff
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    address_id UUID REFERENCES address(id) ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX person_username_idx ON person(username);
CREATE INDEX person_address_idx ON person(address_id);


CREATE TABLE IF NOT EXISTS contact(
    id UUID NOT NULL PRIMARY KEY,
    person_id UUID NOT NULL REFERENCES person(id) ON UPDATE CASCADE ON DELETE CASCADE,
    contact_type TEXT CHECK ( contact_type in ('email', 'home', 'work', 'mobile' )),
    contact_value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX contact_type_idx ON contact(contact_type);

CREATE TABLE IF NOT EXISTS property (
    id UUID NOT NULL PRIMARY KEY,
    property_name TEXT NOT NULL,
    address_id UUID NOT NULL REFERENCES address(id) ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS person_property(
    person_id UUID REFERENCES person(id) ON UPDATE CASCADE,
    property_id UUID REFERENCES property(id) ON UPDATE CASCADE,
    property_relation TEXT CHECK ( property_relation in ('owner', 'renter') ),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT renter_pk PRIMARY KEY(person_id, property_id)
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

CREATE TRIGGER update_address_modified_at_column BEFORE INSERT OR UPDATE ON address FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_contact_modified_at_column BEFORE INSERT OR UPDATE ON contact FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_person_modified_at_column BEFORE INSERT OR UPDATE ON person FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_property_modified_at_column BEFORE INSERT OR UPDATE ON property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_person_property_modified_at_column BEFORE INSERT OR UPDATE ON person_property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_address_modified_at_column ON address;
DROP TRIGGER IF EXISTS update_contact_modified_at_column ON contact;
DROP TRIGGER IF EXISTS update_person_property_modified_at_column ON person_property;
DROP TRIGGER IF EXISTS update_person_modified_at_column on person;
DROP TRIGGER IF EXISTS update_property_modified_at_column ON property;
DROP FUNCTION IF EXISTS update_modified_at_column();

DROP INDEX person_username_idx;
DROP INDEX person_address_idx;
DROP INDEX contact_type_idx;

DROP TABLE address;
DROP TABLE person;
DROP TABLE contact;
DROP TABLE property;
DROP TABLE person_property;