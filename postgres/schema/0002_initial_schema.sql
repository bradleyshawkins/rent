-- +goose Up
CREATE TABLE IF NOT EXISTS address(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    address_1 TEXT NOT NULL,
    address_2 TEXT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zipcode VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE landlord (
    id UUID NOT NULL PRIMARY KEY,
    person_id UUID NOT NULL REFERENCES person(id) ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE renter (
    id UUID NOT NULL PRIMARY KEY,
    person_id UUID NOT NULL REFERENCES person(id) ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

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

CREATE INDEX user_property_property_relation_idx ON person_property(property_relation);

-- function created in previous sql file
CREATE TRIGGER update_address_modified_at_column BEFORE INSERT OR UPDATE ON address FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_landlord_modified_at_column BEFORE INSERT OR UPDATE ON landlord FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_renter_modified_at_column BEFORE INSERT OR UPDATE ON renter FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_property_modified_at_column BEFORE INSERT OR UPDATE ON property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_person_property_modified_at_column BEFORE INSERT OR UPDATE ON person_property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_address_modified_at_column ON address;
DROP TRIGGER IF EXISTS update_person_property_modified_at_column ON person_property;
DROP TRIGGER IF EXISTS update_renter_modified_at_column on renter;
DROP TRIGGER IF EXISTS update_property_modified_at_column ON property;
DROP TRIGGER IF EXISTS update_landlord_modified_at_column ON landlord;

DROP TABLE address;
DROP TABLE landlord;
DROP TABLE renter;
DROP TABLE property;
DROP TABLE person_property;