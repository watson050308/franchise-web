CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(100),
    role INT
);

CREATE TABLE IF NOT EXISTS roles (
  name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS companies (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  name VARCHAR(100) NOT NULL,
  tel_number VARCHAR(20) NOT NULL,
  business_id VARCHAR(256) NOT NULL,
  responsible_person VARCHAR(256) NOT NULL,
  responsible_person_phone VARCHAR(20) NOT NULL,
  primary_contact_phone VARCHAR(20) NOT NULL,
  company_email VARCHAR(256) NOT NULL,
  company_address VARCHAR(256) NOT NULL,
  web_url varchar(1024),
  brand_id INT,
  user_id INT
);

CREATE TABLE IF NOT EXISTS brands (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  name VARCHAR(256) NOT NULL,
  eng_name VARCHAR(256) NOT NULL,
  industry_id INT NOT NULL,
  franchise_info_id INT NOT NULL,
  company_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS industries (
  name VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS franchise_infos (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  service_type VARCHAR(256) NOT NULL,
  franchise_funds INT,
  north_store_count INT,
  south_store_count INT,
  east_store_count INT,
  west_store_count INT,
  central_store_count INT,
  foreign_store_count INT,
  contact_id INT
);

CREATE TABLE IF NOT EXISTS franchise_tag_link (
  franchise_info_id INT,
  search_tag_id INT
);

CREATE TABLE IF NOT EXISTS search_tags (
  name VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS contacts (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  exhibitor_name VARCHAR(100) NOT NULL,
  exhibitor_position VARCHAR(100),
  exhibitor_title VARCHAR(256) NOT NULL,
  phone VARCHAR(20) NOT NULL
);
