
test:
	@go test -v -cover ./...
.PHONY: test

bench:
	@go test -cpu 1,2,4 -bench=. -benchmem
.PHONY: bench

race:
	@go test -cover -cpu 1,4 -bench=. -race ./...
.PHONY: race