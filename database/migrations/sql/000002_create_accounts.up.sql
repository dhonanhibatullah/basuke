CREATE TYPE basuke_account_role AS ENUM (
    'super',
    'admin',
    'user'
);

CREATE TYPE basuke_account_status AS ENUM (
    'active',
    'inactive',
    'suspended'
);

CREATE TABLE basuke_accounts (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role basuke_account_role NOT NULL DEFAULT 'user',
    status basuke_account_status NOT NULL DEFAULT 'active',
    bio VARCHAR(2048),
    preferences JSONB,
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ
);

CREATE INDEX idx_basuke_accounts_username_search ON basuke_accounts USING GIN (username gin_trgm_ops);

CREATE INDEX idx_basuke_accounts_email_search ON basuke_accounts USING GIN (email gin_trgm_ops);

CREATE INDEX idx_basuke_accounts_name_search ON basuke_accounts USING GIN (name gin_trgm_ops);

CREATE INDEX idx_basuke_accounts_role ON basuke_accounts (role);

CREATE INDEX idx_basuke_accounts_status ON basuke_accounts (status);