test:
	go test . -v -cover -race -coverprofile coverage.out

cover:
	go tool cover -html coverage.out -o coverage.html \
	&& ./generate-coverage-badge.sh

.PHONY: all clean test cover
