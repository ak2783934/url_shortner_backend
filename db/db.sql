-- Tables

-- url_shortner
-- ____________________________________________
-- id         | uuid/string
-- short_url  | string (indexed based on this)
-- long_url   | string
-- created_at | timestamp
-- updated_at | timestamp


CREATE TABLE url_shortner (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    short_url VARCHAR(255) NOT NULL,
    long_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);