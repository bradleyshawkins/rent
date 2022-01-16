-- +goose Up

CREATE TABLE account(
    id UUID NOT NULL PRIMARY KEY,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_account_modified_at_column BEFORE INSERT OR UPDATE ON account FOR EACH ROW EXECUTE FUNCTION update_modified_at_column();

CREATE TABLE membership (
    account_id UUID NOT NULL REFERENCES account(id),
    app_user_id UUID NOT NULL REFERENCES app_user(id),
    role_id TEXT NOT NULL,
    PRIMARY KEY (account_id, app_user_id),
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