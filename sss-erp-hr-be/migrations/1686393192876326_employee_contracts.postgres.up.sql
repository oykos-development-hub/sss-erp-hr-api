DROP TABLE if exists employee_contracts;

CREATE TABLE employee_contracts (
    id serial PRIMARY KEY,
    user_profile_id INT,
    contract_type_id INT,
    abbreviation text,
    description TEXT,
    active BOOLEAN NOT NULL,
    serial_number text,
    net_salary text,
    gross_salary text,
    bank_account text,
    bank_name text,
    date_of_signature DATE,
    date_of_eligibility DATE,
    date_of_start DATE NOT NULL,
    date_of_end DATE,
    file_id INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE
);
