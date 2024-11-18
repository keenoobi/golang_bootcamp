

all: build

build: main.go
	@go build main.go

run: build
	@./main < input.txt

test:
	@go test project/parser project/stats     

cover:
	@go test project/parser project/stats -coverprofile=coverage.out -v
	@go tool cover -html=coverage.out -o coverage.html && open coverage.html

clean:
	@rm -rf coverage.* main
	@go clean -testcache
	@go clean -modcache