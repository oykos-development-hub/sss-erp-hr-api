DROP TABLE if exists employee_contracts;

CREATE TABLE employee_contracts (
    id serial PRIMARY KEY,
    user_profile_id INT NOT NULL,
    contract_type_id INT NOT NULL,
    organization_unit_id INTEGER NOT NULL,
    organization_unit_department_id INTEGER,
    job_position_in_organization_unit INTEGER NOT NULL,
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
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (organization_unit_id) REFERENCES organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (organization_unit_department_id) REFERENCES organization_units  (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (job_position_in_organization_unit) REFERENCES job_positions_in_organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE
);
