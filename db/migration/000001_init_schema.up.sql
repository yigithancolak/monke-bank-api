CREATE TABLE currencies (
    code VARCHAR(3) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    symbol VARCHAR(10)
);

INSERT INTO currencies (code, name, symbol)
VALUES 
    ('USD', 'United States Dollar', '$'),
    ('EUR', 'Euro', '€'),
    ('GBP', 'British Pound Sterling', '£'),
    ('JPY', 'Japanese Yen', '¥'),
    ('CNY', 'Chinese Yuan', '¥'),
    ('INR', 'Indian Rupee', '₹'),
    ('RUB', 'Russian Ruble', '₽'),
    ('TRY', 'Turkish Lira', '₺');

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    owner UUID NOT NULL,
    balance INTEGER NOT NULL DEFAULT 0,
    currency_code VARCHAR(3) NOT NULL REFERENCES currencies(code),
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entries (
    id UUID PRIMARY KEY,
    account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfers (
    id UUID PRIMARY KEY,
    from_account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    to_account_id UUID NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);
