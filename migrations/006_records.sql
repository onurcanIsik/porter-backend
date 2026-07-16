

CREATE TABLE IF NOT EXISTS records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    resource_id UUID NOT NULL,
    record_data JSONB NOT NULL,
    record_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    record_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (resource_id) REFERENCES resources(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_records_resource_id ON records(resource_id);