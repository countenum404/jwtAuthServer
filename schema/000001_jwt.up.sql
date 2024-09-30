SELECT gen_random_uuid();

CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fisrtname VARCHAR(255) not null,
    lastname VARCHAR(255) not null,
    username VARCHAR(255) not null,
    email VARCHAR(255) not null,
    pwd VARCHAR(255) not null
);

CREATE TABLE IF NOT EXISTS tokens
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    refresh_token TEXT not null,
    session_id UUID,
    user_id UUID REFERENCES users(id)
);

INSERT INTO users(fisrtname, lastname, username, email, pwd) VALUES
('Terry', 'Davis', 'tdavis', 'd.shabashovqa@gmail.com','c3VwZXJwYXNzd29yZA==');