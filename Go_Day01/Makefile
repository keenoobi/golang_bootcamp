# Переменные для удобства
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=readDB
BINARY_COMPARE_DB=compareDB
BINARY_COMPARE_FS=compareFS
PROJECT=go list ./... | grep -v /cmd/

# Цель по умолчанию
all: build test

# Сборка всех бинарных файлов
build: build-readDB build-compareDB build-compareFS

# Сборка readDB
build-readDB:
	@$(GOBUILD) -o bin/$(BINARY_NAME) cmd/readDB/main.go

# Сборка compareDB
build-compareDB:
	@$(GOBUILD) -o bin/$(BINARY_COMPARE_DB) cmd/compareDB/main.go

# Сборка compareFS
build-compareFS:
	@$(GOBUILD) -o bin/$(BINARY_COMPARE_FS) cmd/compareFS/main.go

# Очистка всех бинарных файлов
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f bin/$(BINARY_COMPARE_DB)
	rm -f bin/$(BINARY_COMPARE_FS)
	rm -f coverage.out coverage.html

# Запуск тестов
test:
	$(GOTEST) -v ./...

# Запуск readDB с аргументами
run-readDB: build-readDB
	@./bin/$(BINARY_NAME) -f testdata/original_database.xml
	@./bin/$(BINARY_NAME) -f testdata/stolen_database.json

# Запуск compareDB с аргументами
run-compareDB: build-compareDB
	@./bin/$(BINARY_COMPARE_DB) --old testdata/original_database.xml --new testdata/stolen_database.json

# Запуск compareFS с аргументами
run-compareFS: build-compareFS
	@./bin/$(BINARY_COMPARE_FS) --old testdata/snapshot1.txt --new testdata/snapshot2.txt
	@./bin/$(BINARY_COMPARE_FS) --old testdata/snapshot3.txt --new testdata/snapshot4.txt

# Установка зависимостей
deps:
	$(GOGET) -v ./...

# Цель для генерации отчета о покрытии тестами
coverage:
	$(GOTEST) -coverprofile=coverage.out `$(PROJECT)`
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	open coverage.html

# Цель для очистки файлов, созданных при генерации покрытия
clean-coverage:
	rm -f coverage.out coverage.html

.PHONY: all build clean test run-readDB run-compareDB run-compareFS deps coverage clean-coverage