run-all:
	go run ./cmd/app/...

mockery-install:
	go install github.com/vektra/mockery/v3@v3.2.5

.PHONY: test
test:
	go test -v -cover ./...

integration-test:
	go test -count=1 -v -tags=integration ./test/integration
sqlc-install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

sqlc-generate:
	sqlc generate -f internal/adapter/postgres/sqlc.yaml