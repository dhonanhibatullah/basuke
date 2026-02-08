CREATE TYPE basuke_auth_status AS ENUM (
    'activated',
    'deactivated',
    'suspended'
);

CREATE TABLE basuke_auths (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL CONSTRAINT fk_auth_user REFERENCES basuke_users (id) ON DELETE CASCADE,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    password_hash VARCHAR(255) NOT NULL,
    status basuke_auth_status NOT NULL DEFAULT 'activated',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT REFERENCES basuke_users (id) ON DELETE SET NULL,
    updated_at TIMESTAMPTZ,
    updated_by BIGINT REFERENCES basuke_users (id) ON DELETE SET NULL,
    CONSTRAINT uk_auths_username UNIQUE (username),
    CONSTRAINT uk_auths_email UNIQUE (email)
);

CREATE INDEX idx_basuke_auths_status ON basuke_auths (status);

CREATE INDEX idx_basuke_auths_user_id ON basuke_auths (user_id);