ALTER TABLE franchisees_info
DROP CONSTRAINT IF EXISTS fk_contact,
DROP COLUMN IF EXISTS contact_id;