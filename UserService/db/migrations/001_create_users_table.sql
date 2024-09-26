CREATE TABLE IF NOT EXISTS users (
                                     user_id UUID PRIMARY KEY,
                                     username TEXT NOT NULL,
                                     email TEXT NOT NULL
);
