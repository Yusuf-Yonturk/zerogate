# ZeroGate Idempotent Payment API

[Türkçe Belge İçin Tıklayın (Read in Turkish)](README.tr.md)

ZeroGate is a distributed payment gateway component that demonstrates how to prevent double-spending in financial transactions using Redis-based distributed locks and the idempotency pattern.

## Architecture
The system uses PostgreSQL for persistence and Redis for acquiring distributed locks before processing requests. This ensures that concurrent duplicate requests are either queued or rejected, maintaining data integrity.

## How to Install and Run

Follow these steps to run the project on your local machine:

1. Start Database and Cache Servers:
   Run PostgreSQL and Redis in the background using Docker.
   ```bash
   docker-compose up -d
   ```

2. Download Dependencies:
   ```bash
   go mod tidy
   ```

3. Start the Application:
   ```bash
   go run cmd/api/main.go
   ```
   When the application starts successfully, you will see the `Server listening on port 8080` message.

## How to Test (Proof of Idempotency)

While the application is running, open a NEW terminal window and follow these steps.

1. **First Request (Normal Payment):**
   Send a payment request by running the following command.
   ```bash
   curl -X POST http://localhost:8080/api/pay \
        -H "Content-Type: application/json" \
        -d '{"idempotency_key": "CUSTOMER-PAY-001", "amount": 250.00}'
   ```
   *Expected Result:* The transaction will be approved, a new ID will be generated, and `status: completed` will be returned.

2. **Second Request (Double-Spending Prevention Test):**
   Run the EXACT SAME command again (simulating the user clicking the pay button a second time).
   ```bash
   curl -X POST http://localhost:8080/api/pay \
        -H "Content-Type: application/json" \
        -d '{"idempotency_key": "CUSTOMER-PAY-001", "amount": 250.00}'
   ```
   *Expected Result:* The system will not charge the money a second time. It will return the EXACT SAME ID generated in the first transaction, indicating that the request was already processed.

## Endpoints
POST /api/pay
Accepts an idempotency key and amount. Guarantees that only one transaction will be processed per key.
