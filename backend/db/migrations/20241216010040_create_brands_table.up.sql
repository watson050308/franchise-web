CREATE TABLE brands (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    brand_name VARCHAR(256) NOT NULL,
    brand_eng_name VARCHAR(256) NOT NULL,
    industry_id INTEGER NOT NULL,
    franchise_info_id INTEGER NOT NULL
);