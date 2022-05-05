BEGIN;

DROP INDEX IF EXISTS audit_object_type_object_id_idx;

DROP TABLE IF EXISTS audit;

COMMIT;