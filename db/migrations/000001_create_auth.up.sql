CREATE TABLE IF NOT EXISTS users
(
    id         serial PRIMARY KEY,
    first_name varchar(32),
    last_name  varchar(32)
);

CREATE TABLE IF NOT EXISTS auth_passwords
(
    email    varchar(255) PRIMARY KEY,
    password VARCHAR(64) NOT NULL,
    user_id  int NOT NULL REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);