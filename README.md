# ZeroGate Idempotent Payment API

[Türkçe Belge İçin Tıklayın (Read in Turkish)](README.tr.md)

ZeroGate is a distributed payment gateway component that demonstrates how to prevent double-spending in financial transactions using Redis-based distributed locks and the idempotency pattern.

## Architecture
The system uses PostgreSQL for persistence and Redis for acquiring distributed locks before processing requests. This ensures that concurrent duplicate requests are either queued or rejected, maintaining data integrity.

## Endpoints
POST /api/pay
Accepts an idempotency key and amount. Guarantees that only one transaction will be processed per key.
