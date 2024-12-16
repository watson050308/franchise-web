CREATE TABLE franchise_fees_info(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    paid_url VARCHAR(100),
    paid_fee INTEGER NOT NULL DEFAULT 0,
    franchise_info_id INTEGER NOT NULL,
    CONSTRAINT franchise_fee_info FOREIGN KEY (franchise_info_id) REFERENCES franchisees_info (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);