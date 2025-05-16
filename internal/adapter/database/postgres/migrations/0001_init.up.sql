CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firebase_uuid TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    security_image TEXT,
    security_phrase TEXT
);

CREATE INDEX idx_users_firebase_uuid ON users(firebase_uuid);