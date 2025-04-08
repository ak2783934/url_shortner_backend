-- Create the database
CREATE DATABASE IF NOT EXISTS url_shortener;

-- Use the database
USE url_shortener;

-- Create a table to store short URLs
CREATE TABLE IF NOT EXISTS url_shortner (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    short_url VARCHAR(255) NOT NULL,
    long_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);