
test:
	@go test -v -cover ./...
.PHONY: test

bench:
	@go test -cover -cpu 1,4 -bench=. ./...
.PHONY: bench

race:
	@go test -cover -cpu 1,4 -bench=. -race ./...
.PHONY: race