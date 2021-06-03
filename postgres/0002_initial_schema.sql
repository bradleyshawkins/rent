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

CREATE TABLE user (
    id UUID NOT NULL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL, -- todo add better password stuff
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email_address TEXT NOT NULL,
    address_id UUID REFERENCES address(id) ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_username_idx ON user(username);


CREATE TABLE IF NOT EXISTS property (
    id UUID NOT NULL PRIMARY KEY,
    property_name TEXT NOT NULL,
    address_id UUID NOT NULL REFERENCES address(id) ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_property(
    user_id UUID REFERENCES user(id) ON UPDATE CASCADE,
    property_id UUID REFERENCES property(id) ON UPDATE CASCADE,
    property_relation TEXT CHECK ( property_relation in ('owner', 'renter') ),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT renter_pk PRIMARY KEY(user_id, property_id)
);

CREATE UNIQUE INDEX user_property_user_property_idx ON user_property(user_id, property_id);
CREATE INDEX user_property_property_relation_idx ON user_property(property_relation);

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
CREATE TRIGGER update_user_modified_at_column BEFORE INSERT OR UPDATE ON user FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_property_modified_at_column BEFORE INSERT OR UPDATE ON property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_user_property_modified_at_column BEFORE INSERT OR UPDATE ON user_property FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_address_modified_at_column ON address;
DROP TRIGGER IF EXISTS update_user_property_modified_at_column ON user_property;
DROP TRIGGER IF EXISTS update_user_modified_at_column on user;
DROP TRIGGER IF EXISTS update_property_modified_at_column ON property;
DROP FUNCTION IF EXISTS update_modified_at_column();

DROP INDEX user_username_idx;

DROP TABLE address;
DROP TABLE user;
DROP TABLE property;
DROP TABLE user_property;