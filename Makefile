.PHONY: build test lint format check clean install

install:
	cd web && npm install
	cd api && go mod download

build: build-web build-api

build-web:
	cd web && npx next build

build-api:
	cd api && go build -trimpath -ldflags "-s -w" -o bin/server ./cmd/server

test: test-web test-api

test-web:
	cd web && npx vitest run --passWithNoTests

test-api:
	cd api && go test -v -race -count=1 ./...

lint: lint-web lint-api

lint-web:
	cd web && npx oxlint .
	cd web && npx biome check .

lint-api:
	cd api && golangci-lint run ./...

format:
	cd web && npx biome format --write .
	cd api && gofumpt -w . && goimports -w .

check: format lint test build
	@echo "All checks passed."

clean:
	rm -rf web/.next web/node_modules/.cache api/bin/
