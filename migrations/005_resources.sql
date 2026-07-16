

CREATE TABLE IF NOT EXISTS resources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL,
    resource_name VARCHAR(255) NOT NULL,
    resource_schema JSONB NOT NULL,
    resource_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    resource_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
    UNIQUE (project_id, resource_name)
);