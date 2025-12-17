SHELL := /bin/bash

.PHONY: help
help:
	@echo "Targets:"
	@echo "  up          - run local stack via docker compose"
	@echo "  down        - stop local stack"
	@echo "  test        - run unit tests (Go services)"
	@echo "  build       - build Docker images"
	@echo "  k8s-apply   - apply all Kubernetes manifests"
	@echo "  k8s-delete  - delete all Kubernetes manifests"
	@echo "  k8s-apply-minimal   - apply core K8s (deploy + svc)"
	@echo "  k8s-delete-minimal  - delete core K8s (deploy + svc)"
	@echo "  proto       - generate protobuf code (requires protoc + plugins)"

.PHONY: up
up:
	docker compose -f infra/docker-compose.yml up -d --build

.PHONY: down
down:
	docker compose -f infra/docker-compose.yml down -v

.PHONY: test
test:
	( cd services/catalog-go && go test ./... )
	( cd services/orders-go && go test ./... )

.PHONY: build
build:
	docker build -t mszth/catalog:dev services/catalog-go
	docker build -t mszth/orders:dev services/orders-go

.PHONY: k8s-apply
k8s-apply:
	kubectl apply -f k8s/

.PHONY: k8s-delete
k8s-delete:
	kubectl delete -f k8s/ --ignore-not-found=true

.PHONY: k8s-apply-minimal
k8s-apply-minimal:
	kubectl apply -f k8s/00-namespace.yaml
	kubectl apply -f k8s/catalog-svc.yaml -f k8s/catalog-deploy.yaml
	kubectl apply -f k8s/orders-svc.yaml -f k8s/orders-deploy.yaml

.PHONY: k8s-delete-minimal
k8s-delete-minimal:
	kubectl delete -f k8s/orders-deploy.yaml -f k8s/orders-svc.yaml --ignore-not-found=true
	kubectl delete -f k8s/catalog-deploy.yaml -f k8s/catalog-svc.yaml --ignore-not-found=true
	kubectl delete -f k8s/00-namespace.yaml --ignore-not-found=true

.PHONY: proto
proto:
	@echo "Run protoc to generate Go code from proto/orders.proto (see proto/README.md)"
