CREATE TABLE users(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username varchar(256) NOT NULL UNIQUE,
    email varchar(256) UNIQUE,
    password varchar(256) NOT NULL,
    name varchar(256)
);

CREATE TABLE notes(
    id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES notes(id) 
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    user_id uuid REFERENCES users(id)
        ON UPDATE CASCADE 
        ON DELETE CASCADE,
    note TEXT,
    done INTEGER NOT NULL
);