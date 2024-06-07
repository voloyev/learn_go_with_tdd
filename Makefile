.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: test/bench
test/bench:
	go test -v -race -buildvcs -bench=. ./...

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -bench=. -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out
