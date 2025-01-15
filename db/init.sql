CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    amount DOUBLE PRECISION NOT NULL,
    type VARCHAR(50) NOT NULL,
    parent_id BIGINT REFERENCES transactions(id)
);
