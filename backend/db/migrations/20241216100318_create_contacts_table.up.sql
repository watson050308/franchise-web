CREATE TABLE contacts(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    exhibitor_name VARCHAR(100) NOT NULL,
    exhibitor_title VARCHAR(100) NOT NULL,
    exhibitor_email VARCHAR(256) NOT NULL,
    phone_number VARCHAR(20) NOT NULL
);