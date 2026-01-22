-- case-insensitive email
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
                                     id UUID DEFAULT uuidv7() PRIMARY KEY,

                                     email CITEXT NOT NULL,
                                     login TEXT NOT NULL,
                                     password_hash TEXT NOT NULL,

                                     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                                     updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

                                     email_verified_at TIMESTAMPTZ NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email_uq ON users (email);
CREATE UNIQUE INDEX IF NOT EXISTS users_login_uq ON users (login);