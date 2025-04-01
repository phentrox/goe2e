-- +dbevo Up
CREATE TABLE example (
    id BIGINT UNIQUE PRIMARY KEY,
    name TEXT,
    username TEXT
);

CREATE TABLE example3 (
    id BIGINT UNIQUE PRIMARY KEY,
    name TEXT,
    username TEXT
);
