CREATE TABLE IF NOT EXISTS users(
  id uuid PRIMARY KEY,
  username varchar UNIQUE NOT NULL,
  email varchar UNIQUE NOT NULL,
  password varchar NOT NULL,
  updated_at timestamp,
  created_at timestamp NOT NULL DEFAULT now()
);