CREATE TABLE IF NOT EXISTS tickets (
    id INTEGER PRIMARY KEY NOT NULL,
    name TEXT,
    price REAL,
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);