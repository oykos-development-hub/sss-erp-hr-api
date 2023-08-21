CREATE TABLE IF NOT EXISTS tender_applications_in_organization_units (
    id serial PRIMARY KEY,
    job_tender_id int NOT NULL,
    user_profile_id int,
    active boolean not null,
    is_internal boolean not null,
    first_name text,
    last_name text,
    official_personal_id text,
    nationality text,
    evaluation text,
    date_of_birth date,
    date_of_application date,
    status text,
    file_id int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (job_tender_id) REFERENCES tenders_in_organization_units (id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_profile_id) REFERENCES user_profiles (id) ON UPDATE CASCADE ON DELETE CASCADE
);