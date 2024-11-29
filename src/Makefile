all: build

build:
	@go build -o bin/ -v ./cmd/anscombe 

test:
	@go test -v project/tests

cover:
	@go test -v project/tests -coverprofile=coverage.out 
	@go tool cover -html=coverage.out -o coverage.html && open coverage.html

clean:
	@rm -rf coverage.* main
	@go clean -testcache
	@go clean -modcache