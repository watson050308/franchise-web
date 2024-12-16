ALTER TABLE franchisees_info
ADD CONSTRAINT fk_contact FOREIGN KEY (contact_id) REFERENCES contacts (id)
    ON UPDATE CASCADE
    ON DELETE SET NULL; 