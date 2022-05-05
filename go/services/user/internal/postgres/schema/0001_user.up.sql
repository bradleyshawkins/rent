BEGIN;

CREATE TABLE IF NOT EXISTS app_user (
    id UUID NOT NULL PRIMARY KEY,
    status TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email_address TEXT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS app_user_status_idx ON app_user(status);
CREATE UNIQUE INDEX IF NOT EXISTS app_user_email_address ON app_user(email_address);
CREATE UNIQUE INDEX IF NOT EXISTS app_user_username ON app_user(username);

CREATE OR REPLACE FUNCTION update_modified_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at = NOW();
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER update_app_user_modified_at_column BEFORE INSERT OR UPDATE ON app_user FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

COMMIT;
