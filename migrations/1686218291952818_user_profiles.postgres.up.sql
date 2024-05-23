DROP TABLE if exists user_profiles;

CREATE TABLE user_profiles (
    id serial PRIMARY KEY,
    user_account_id int NOT NULL,
    first_name text NOT NULL,
    middle_name text,
    last_name text NOT NULL,
    birth_last_name text,
    father_name text NOT NULL,
    mother_name text NOT NULL,
    mother_birth_last_name text,
    date_of_birth date NOT NULL,
    country_of_birth text NOT NULL,
    city_of_birth text NOT NULL,
    nationality text,
    national_minority text,
    citizenship text NOT NULL,
    address text NOT NULL,
    bank_account text,
    bank_name text,
    official_personal_id text NOT NULL,
    official_personal_document_number text NOT NULL,
    official_personal_document_issuer text NOT NULL,
    gender text NOT NULL,
    single_parent boolean NOT NULL,
    housing_done boolean NOT NULL,
    housing_description text NOT NULL,
    martial_status text NOT NULL,
    date_of_taking_oath date,
    date_of_becoming_judge text,
    judge_application_submission_date text,
    active_contract boolean,
    revisor_role boolean,
    engagement_type_id integer,
    is_judge bool,
    personal_id text,
    year integer,
    file_id integer not null,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now()
);

INSERT INTO user_profiles (
    user_account_id, first_name, middle_name, last_name, birth_last_name, father_name, mother_name, mother_birth_last_name,
    date_of_birth, country_of_birth, city_of_birth, nationality, national_minority, citizenship, address, bank_account, 
    bank_name, official_personal_id, official_personal_document_number, official_personal_document_issuer, gender, 
    single_parent, housing_done, housing_description, martial_status, date_of_taking_oath, date_of_becoming_judge, 
    active_contract, revisor_role, engagement_type_id, created_at, updated_at, personal_id, is_judge, file_id, year, 
    judge_application_submission_date
) VALUES (
    1, 'Nikola', '', 'Perović', '', 'Marko', 'Mila', '', '1990-05-05', 'mne', 'Podgorica', '', '', 'mne', 
    'Ulica Vladana Giljena 21', '', '', '0505990191919', 'DOC67890', 'PJ Bar', 'M', 'f', 't', 'Stan u Nikšiću', 
    'Oženjen', '2023-10-26', NULL, NULL, NULL, NULL, '2023-10-26 08:30:48.534364', '2023-10-26 08:30:48.534366', '', 'f', 
    0, 2024, NULL
);