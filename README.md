# Transaction Service

A RESTful service that manages financial transactions with parent-child relationships and provides sum calculations across transaction hierarchies.

## Features

- Create and update transactions
- Retrieve transaction details
- Get transactions by type
- Calculate sum of linked transactions
- Prevent cyclic transaction relationships
- Structured JSON logging
- Configuration management using TOML

## API Endpoints

### PUT /transactionservice/transaction/{transaction_id}
Create or update a transaction

### GET /transactionservice/transaction/{transaction_id}
Retrieve transaction details

### GET /transactionservice/types/{type}
Get all transaction IDs of a specific type

### GET /transactionservice/sum/{transaction_id}
Calculate the sum of the specified transaction and all its children

## Setup

1. Clone the repository
2. Copy the configuration example to file name `config/credentials.toml`
3. Update the database credentials in `config/credentials.toml`
4. Run the docker compose file to start the database on localhost:5432
```bash
docker compose up -d
```
5. Build and run the service:
```bash
go run main.go
```

