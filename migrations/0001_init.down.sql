-- Drop tables in reverse order

-- to avoid foreign key constraint violations
DROP TABLE IF EXISTS patients;
DROP TABLE IF EXISTS users;
