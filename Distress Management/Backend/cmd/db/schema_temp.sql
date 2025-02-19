-- Create database
CREATE DATABASE IF NOT EXISTS distress_management;
USE distress_management;

-- Disable foreign key checks
SET FOREIGN_KEY_CHECKS = 0;

-- Drop existing tables if they exist
DROP TABLE IF EXISTS documents;
DROP TABLE IF EXISTS progress_notes;
DROP TABLE IF EXISTS cases;

-- Enable foreign key checks
SET FOREIGN_KEY_CHECKS = 1;

-- Cases table
CREATE TABLE IF NOT EXISTS cases (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    reference_number VARCHAR(50) NOT NULL UNIQUE,
    sender_name VARCHAR(255) NOT NULL,
    receiving_date TIMESTAMP NOT NULL,
    subject VARCHAR(255) NOT NULL,
    country_of_origin VARCHAR(100) NOT NULL,
    distressed_person_name VARCHAR(255) NOT NULL,
    nature_of_case ENUM('Emergency', 'Urgent', 'Standard') NOT NULL,
    case_details TEXT NOT NULL,
    status ENUM('Pending', 'Under Review', 'Assigned', 'In Progress', 'Resolved', 'Closed') NOT NULL DEFAULT 'Pending',
    stage ENUM('Front Office Receipt', 'Director Review', 'Cadet Assignment', 'Case Investigation', 'Case Resolution') NOT NULL DEFAULT 'Front Office Receipt',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Documents table
CREATE TABLE IF NOT EXISTS documents (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    case_id BIGINT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (case_id) REFERENCES cases(id) ON DELETE CASCADE
);

-- Progress notes table
CREATE TABLE IF NOT EXISTS progress_notes (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    case_id BIGINT NOT NULL,
    note TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (case_id) REFERENCES cases(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_cases_reference_number ON cases(reference_number);
CREATE INDEX idx_cases_status ON cases(status);
CREATE INDEX idx_cases_stage ON cases(stage);
CREATE INDEX idx_progress_notes_case_id ON progress_notes(case_id);
