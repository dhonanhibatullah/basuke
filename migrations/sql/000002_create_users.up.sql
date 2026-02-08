CREATE TYPE basuke_user_role AS ENUM (
    'super',
    'admin',
    'client'
);

CREATE TYPE basuke_user_status AS ENUM (
    'online',
    'offline',
    'away',
    'dnd'
);

CREATE TABLE basuke_users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    role basuke_user_role NOT NULL DEFAULT 'client',
    status basuke_user_status NOT NULL DEFAULT 'offline',
    bio VARCHAR(2048),
    last_login_at TIMESTAMPTZ,
    preferences JSONB,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT REFERENCES basuke_users (id) ON DELETE SET NULL,
    updated_at TIMESTAMPTZ,
    updated_by BIGINT REFERENCES basuke_users (id) ON DELETE SET NULL
);

CREATE INDEX idx_basuke_users_role ON basuke_users (role);

CREATE INDEX idx_basuke_users_status ON basuke_users (status);

CREATE INDEX idx_basuke_users_search ON basuke_users USING GIN (name gin_trgm_ops);