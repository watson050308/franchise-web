CREATE TABLE franchisees_info (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    tel_number VARCHAR(20) NOT NULL,
    service_type VARCHAR(256) NOT NULL,
    franchise_funds INTEGER NOT NULL,
    north_store_count INTEGER NOT NULL DEFAULT 0,
    south_store_count INTEGER NOT NULL DEFAULT 0,
    east_store_count INTEGER NOT NULL DEFAULT 0,
    central_store_count INTEGER NOT NULL DEFAULT 0,
    foreign_store_count INTEGER NOT NULL DEFAULT 0,
    contact_id INTEGER NOT NULL
);