CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    hashed_password bytea NOT NULL,
    activated bool  Not NULL,
    version integer Not Null DEFAULT 1
);