

CREATE TABLE IF NOT EXISTS users (
    id UUID  PRIMARY KEY DEFAULT gen_random_uuid(),
    provider      VARCHAR(20) NOT NULL,
    provider_id   VARCHAR(255) NOT NULL,
    user_mail VARCHAR(255) NOT NULL UNIQUE,
    user_name VARCHAR(255) NOT NULL,
    user_profile_url VARCHAR(255),
    is_premium BOOLEAN DEFAULT FALSE,
    user_job_title VARCHAR(255) DEFAULT 'Software Engineer',
    user_device_id VARCHAR(255),
    user_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(provider, provider_id)
);