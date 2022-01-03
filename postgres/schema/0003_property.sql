-- +goose Up

CREATE TABLE property(
    id UUID NOT NULL PRIMARY KEY,
    account_id UUID NOT NULL REFERENCES account(id),
    property_status TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX property_property_status_idx ON property(property_status);

CREATE TRIGGER update_property_modified_at_column BEFORE INSERT OR UPDATE ON property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

CREATE TABLE address(
    id UUID NOT NULL PRIMARY KEY,
    street_1 TEXT NOT NULL,
    street_2 TEXT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zipcode TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_address_details_modified_at_column BEFORE INSERT OR UPDATE ON address FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

CREATE TABLE property_address(
    property_id UUID NOT NULL REFERENCES property(id),
    address_id UUID NOT NULL REFERENCES address(id),
    PRIMARY KEY (property_id, address_id)
);

-- +goose Down
DROP TRIGGER update_address_details_modified_at_column ON address;
DROP TRIGGER update_property_modified_at_column ON property;

DROP TABLE property_address;
DROP TABLE address;
DROP TABLE property;