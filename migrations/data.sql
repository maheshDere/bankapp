CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    is_customer BOOLEAN DEFAULT true,
    email VARCHAR(320) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
)


CREATE TABLE accounts(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL FOREIGN KEY REFERENCES users (id) ON DELETE CASCADE,
    account_no VARCHAR(255),
    created_at TIMESTAMP DEFAULT current_timestamp
)

CREATE TABLE transaction_history(
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL FOREIGN KEY REFERENCES accounts (id) ON DELETE CASCADE,
    type VARCHAR(255),
    transaction_amount NUMERIC,
    total_balance NUMERIC,
    transaction_date timestamp NOT NULL
)
