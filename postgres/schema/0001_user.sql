-- +goose Up

CREATE TABLE app_user_details(
    id UUID NOT NULL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE app_user (
    id UUID NOT NULL PRIMARY KEY,
    email_address TEXT NOT NULL,
    password TEXT NOT NULL,
    status TEXT NOT NULL,
    app_user_details_id UUID NOT NULL REFERENCES app_user_details(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX user_username_idx ON app_user(email_address);
CREATE INDEX app_user_status_idx ON app_user(status);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_modified_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at = NOW();
    RETURN NEW;
END;
$$ language plpgsql;
-- +goose StatementEnd

CREATE TRIGGER update_app_user_modified_at_column BEFORE INSERT OR UPDATE ON app_user FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();
CREATE TRIGGER update_app_user_details_modified_at_column BEFORE INSERT OR UPDATE ON app_user_details FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

-- +goose Down
DROP TRIGGER IF EXISTS update_app_user_details_modified_at_column ON app_user_details;
DROP TRIGGER IF EXISTS update_app_user_modified_at_column ON app_user;

DROP FUNCTION update_modified_at_column();

DROP TABLE app_user;
DROP TABLE app_user_details;