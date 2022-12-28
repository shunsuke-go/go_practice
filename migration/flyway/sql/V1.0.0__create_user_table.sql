CREATE TABLE if NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    createdAt timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updatedAt timestamptz NOT NULL default CURRENT_TIMESTAMP
)