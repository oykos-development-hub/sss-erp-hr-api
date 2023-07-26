drop table if exists systematizations;

CREATE TABLE systematizations (
    id serial PRIMARY KEY,
    user_profile_id INT,
    organization_unit_id INTEGER REFERENCES organization_units(id) ON DELETE CASCADE,
    description text NULL,
    serial_number text NOT NULL,
    active BOOL NOT NULL DEFAULT FALSE,
    date_of_activation TIMESTAMP,
    file_id INT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);