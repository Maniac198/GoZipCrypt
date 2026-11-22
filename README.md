# ZipCrypt

ZipCrypt is a learning-focused, multi-tenant secure file exchange API written in Go. It will support streaming file storage, envelope encryption, controlled sharing, asynchronous processing, and audit-ready access control.

Repository: [github.com/Maniac198/GoZipCrypt](https://github.com/Maniac198/GoZipCrypt)

## Current milestone: Stage 0 — foundation

The repository currently provides:

- a dependency-light Go HTTP API with liveness and readiness endpoints;
- validated environment-based configuration;
- structured JSON logs and graceful shutdown;
- Docker Compose services for PostgreSQL, Redis, and MinIO;
- repeatable developer commands through `make`.

## Prerequisites

- Go 1.27+
- Docker Desktop / Docker Compose

## Run locally

```bash
cp .env.example .env
make up
make run
```

`make` loads `.env` only when the file exists. Production environments should
inject these values through their secret and configuration system instead.

In a second terminal:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/readyz
```

MinIO's local console is at http://localhost:9001. Its development credentials are in `.env.example`; never reuse them outside local development.

## Useful commands

```bash
make test
make fmt
make vet
make build
make down
```

## Why there are two health endpoints

`/healthz` answers, “is this process alive?” It should not depend on external systems. `/readyz` answers, “should a load balancer send requests here?” It will check PostgreSQL, Redis, and object storage once those clients exist.

## Next milestone

We will add database migrations, explicit application dependencies, and real readiness checks before implementing tenant identity and authentication.
