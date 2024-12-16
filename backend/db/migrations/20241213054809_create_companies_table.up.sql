CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    company_name VARCHAR(100) NOT NULL,
    tel_number VARCHAR(20) NOT NULL,
    business_id VARCHAR(256) UNIQUE NOT NULL,
    responsible_person VARCHAR(256) NOT NULL,
    responsible_person_phone VARCHAR(20) NOT NULL,
    company_email VARCHAR(256) NOT NULL,
    company_address VARCHAR(256) NOT NULL,
    web_url VARCHAR(1024),
    brand_id INTEGER
);