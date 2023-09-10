CREATE TABLE users (
  id UUID NOT NULL PRIMARY KEY,
  email VARCHAR UNIQUE NOT NULL,
  password VARCHAR NOT NULL,
  full_name VARCHAR NOT NULL,
--   "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE accounts
ADD FOREIGN KEY (owner) REFERENCES users (id);

