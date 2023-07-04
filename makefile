test:
	go test . -v -cover -race -coverprofile cover.out && go tool cover -html=cover.out
.PHONY: all clean test
