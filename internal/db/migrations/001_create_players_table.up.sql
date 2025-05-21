CREATE TABLE IF NOT EXISTS players (
                         id SERIAL PRIMARY KEY,
                         name TEXT NOT NULL,
                         age INT,
                         position TEXT,
                         team TEXT
);