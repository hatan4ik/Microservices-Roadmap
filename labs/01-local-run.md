# Lab 01: Run the reference system locally

From repo root:

```bash
cd reference-implementation
make up
```

Validate:
- Catalog: http://localhost:8080/v1/products
- Orders gRPC: localhost:50051 (use grpcurl if you have it)

Shutdown:
```bash
make down
```
