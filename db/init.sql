CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    amount DOUBLE PRECISION NOT NULL,
    type VARCHAR(50) NOT NULL,
    parent_id BIGINT REFERENCES transactions(id)
);

CREATE INDEX idx_transactions_id ON transactions(id);
CREATE INDEX idx_transactions_parent_id ON transactions(parent_id);
CREATE INDEX idx_transactions_type ON transactions(type);
