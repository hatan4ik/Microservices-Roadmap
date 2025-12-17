# Lab 01: Run the reference system locally

**Goal:** get the local stack running so you can validate endpoints and use it as a sandbox for the chapters.

## Prerequisites

- Docker Desktop (or Docker Engine) with Docker Compose v2
- `make` (optional)

Optional:
- `grpcurl` (to call the gRPC service from your terminal)

## Run it

From repo root:

```bash
cd reference-implementation
make up
```

No `make`? Run Docker Compose directly:

```bash
docker compose -f infra/docker-compose.yml up -d --build
```

This starts:
- `catalog` (HTTP) on `:8080`
- `orders` (gRPC) on `:50051`
- `postgres` on `:5432`
- `redpanda` (Kafka-compatible broker) on `:9092`

## Validate

- Catalog: http://localhost:8080/v1/products
- Health: http://localhost:8080/healthz

Orders gRPC: `localhost:50051` (example using `grpcurl`):

```bash
grpcurl -plaintext -d '{"customer_id":"c1","item_ids":["p1","p2"]}' localhost:50051 orders.v1.OrdersService/CreateOrder
```

## Shutdown / cleanup

```bash
make down
```

No `make`? Run:

```bash
docker compose -f infra/docker-compose.yml down -v
```

## Troubleshooting

- If ports are in use, stop the conflicting container/process or change the published ports in `reference-implementation/infra/docker-compose.yml`.
- If builds fail, run `docker compose -f infra/docker-compose.yml build` from `reference-implementation/` to see detailed output.
