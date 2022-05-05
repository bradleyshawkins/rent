BEGIN;

CREATE TABLE audit (
    id uuid PRIMARY KEY,
    user_id uuid,
    object_type TEXT NOT NULL,
    object_id uuid NOT NULL,
    action TEXT NOT NULL,
    old_value TEXT,
    new_value TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS audit_object_type_object_id_idx ON audit(object_type, object_id);

COMMIT;