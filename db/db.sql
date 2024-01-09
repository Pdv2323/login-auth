-- Create the database
CREATE DATABASE login_auth;
\c login_auth;

-- Enable the uuid-ossp extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    mobile_number VARCHAR(15) NOT NULL,
    password VARCHAR(255) NOT NULL,
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP --NOW()
);

-- Insert a sample user with randomly generated ID
-- INSERT INTO users (first_name, last_name, email, mobile_number, password)
-- VALUES ('John', 'Doe', 'john.doe@example.com', '1234567890', 'password123');



TABLE IF NOT EXISTS users(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    is_staff BOOLEAN DEFAULT FALSE,
    is_superuser BOOLEAN DEFAULT FALSE,
    date_joined TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS users_id_email_is_active_indx ON users (id, email, is_active);
-- -- Create a domain for phone data type
-- CREATE DOMAIN phone AS TEXT CHECK(
--     octet_length(VALUE) BETWEEN 1
--     /*+*/
--     + 8 AND 1
--     /*+*/
--     + 15 + 3
--     AND VALUE ~ '^\+\d+$'
-- );
-- -- User details table (One-to-one relationship)
-- CREATE TABLE user_profile (
--     id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
--     user_id UUID NOT NULL UNIQUE,
--     phone_number phone NULL,
--     birth_date DATE NULL,
--     github_link TEXT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
-- );
-- CREATE INDEX IF NOT EXISTS users_detail_id_user_id ON user_profile (id, user_id);


-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS user_profile;


-- migrate -path=./migrations -database=<DATABASE_URL> up