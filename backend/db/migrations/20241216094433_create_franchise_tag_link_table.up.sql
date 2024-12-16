CREATE TABLE franchise_tag_link (
    franchise_info_id INTEGER NOT NULL,
    search_tag_id INTEGER NOT NULL,
    PRIMARY KEY (franchise_info_id, search_tag_id),
    CONSTRAINT fk_franchise_info FOREIGN KEY (franchise_info_id) REFERENCES franchisees_info (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_search_tag FOREIGN KEY (search_tag_id) REFERENCES search_tags (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);