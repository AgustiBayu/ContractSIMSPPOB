CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    transaction_type VARCHAR(50) NOT NULL
);