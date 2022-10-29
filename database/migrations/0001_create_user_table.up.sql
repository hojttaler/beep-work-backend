CREATE TABLE IF NOT EXISTS "users" (
    id VARCHAR(50) PRIMARY KEY,
    email VARCHAR(80) UNIQUE,
    username VARCHAR(50) UNIQUE,
    role VARCHAR(10),
    status INTEGER,
    password VARCHAR(255),
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE INDEX active_users ON users(id) WHERE status = 1;
