PROJECT=day00
MODULE=project
all: build

build:
	@go build -o bin/ ./cmd/$(PROJECT)
	@echo "the build was successful"

test:
	@go test -v $(MODULE)/tests

cover:
	@go test -v $(MODULE)/internal/... $(MODULE)/tests -coverprofile=coverage.out 
	@go tool cover -html=coverage.out -o coverage.html && open coverage.html

clean:
	@rm -rf coverage.* bin/$(PROJECT)
	@go clean -testcache
	@go clean -modcache
	@echo "everything is clean now"