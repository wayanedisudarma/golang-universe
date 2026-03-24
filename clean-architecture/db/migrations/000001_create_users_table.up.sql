CREATE SCHEMA IF NOT EXISTS clean_architecture;

CREATE TABLE IF NOT EXISTS clean_architecture.users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);