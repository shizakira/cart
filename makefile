mockery-install:
	go install github.com/vektra/mockery/v3@v3.2.5

.PHONY: test
test:
	go test -v -cover ./...
