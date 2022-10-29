CREATE TABLE IF NOT EXISTS "sessions" (
    id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50),
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW(),
    CONSTRAINT fk_sessions_users
        FOREIGN KEY (user_id)
        REFERENCES users(id)
);