

CREATE TABLE IF NOT EXISTS projects (
    id UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    project_name VARCHAR(255) NOT NULL,
    api_key VARCHAR(255) NOT NULL UNIQUE,
    project_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    project_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_projects_user_id ON projects(user_id);