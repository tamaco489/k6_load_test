# =================================================================
# setup tools
# =================================================================
.PHONY: install-tools
install-tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	npm i -g @redocly/cli@latest


# =================================================================
# oapi-codegen
# =================================================================
.PHONY: gen-api bundle-openapi
gen-api: bundle-openapi
	oapi-codegen -package gen -generate types -o internal/gen/types.gen.go ./spec/openapi.yaml
	oapi-codegen -package gen -generate strict-server,gin -templates ./spec/templates -o internal/gen/server.gen.go ./spec/openapi.yaml

bundle-openapi:
	redocly bundle ./spec/openapi_base.yaml --output ./spec/openapi.yaml


# =================================================================
# for CI
# =================================================================
.PHONY: lint deps build run

lint:
	golangci-lint run --timeout 3m

deps:
	go mod download && go mod tidy

build: deps
	go build -o build/jester_api cmd/main.go

run: build
	./build/jester_api


# =================================================================
# for Container
# =================================================================
.PHONY: up down logs
up:
	docker compose up -d jester-api

down:
	docker compose down jester-api

logs:
	docker compose logs -f jester-api