BEGIN;

DROP TRIGGER IF EXISTS update_app_user_modified_at_column ON app_user;

DROP FUNCTION IF EXISTS update_modified_at_column();

DROP TABLE IF EXISTS app_user;

COMMIT;