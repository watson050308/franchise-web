CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    expire_at TIMESTAMP,
    user_name VARCHAR(100) NOT NULL,
    user_email VARCHAR(256) UNIQUE NOT NULL,
    user_password VARCHAR(256) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    role_id INTEGER NOT NULL
);

CREATE INDEX idx_user_email ON users (user_email);
CREATE INDEX idx_expire_at ON users (expire_at);