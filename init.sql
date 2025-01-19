CREATE TABLE IF NOT EXISTS videos (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid()::text,
    title VARCHAR NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    thumbnails VARCHAR
);