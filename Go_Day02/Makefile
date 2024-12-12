# Переменные для удобства
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_FIND=myFind
BINARY_WC=myWc
BINARY_XARGS=myXargs
BINARY_ROTATE=myRotate
PROJECT=go list ./... | grep -v /cmd/

# Цель по умолчанию
all: build test

# Создание папки bin, если она не существует
bin:
	@mkdir -p bin

# Сборка всех бинарных файлов
build: bin build-find build-wc build-xargs build-rotate

# Сборка myFind
build-find: bin
	@$(GOBUILD) -o bin/$(BINARY_FIND) cmd/myFind/main.go

# Сборка myWc
build-wc: bin
	@$(GOBUILD) -o bin/$(BINARY_WC) cmd/myWc/main.go

# Сборка myXargs
build-xargs: bin
	@$(GOBUILD) -o bin/$(BINARY_XARGS) cmd/myXargs/main.go

# Сборка myRotate
build-rotate: bin
	@$(GOBUILD) -o bin/$(BINARY_ROTATE) cmd/myRotate/main.go

# Очистка всех бинарных файлов
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_FIND)
	rm -f bin/$(BINARY_WC)
	rm -f bin/$(BINARY_XARGS)
	rm -f bin/$(BINARY_ROTATE)
	rm -f coverage.out coverage.html

# Запуск тестов
test:
	$(GOTEST) -v ./...

# Запуск myFind с аргументами
run-find: build-find
	@./bin/$(BINARY_FIND) -f -ext 'go' /path/to/dir

# Запуск myWc с аргументами
run-wc: build-wc
	@./bin/$(BINARY_WC) -w input.txt

# Запуск myXargs с аргументами
run-xargs: build-xargs
	@echo "cmd/\ninternal/\n../" | ./bin/$(BINARY_XARGS) ls -la

# Запуск myRotate с аргументами
run-rotate: build-rotate
	@./bin/$(BINARY_ROTATE) /path/to/logs/some_application.log

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

.PHONY: all build clean test run-find run-wc run-xargs run-rotate deps coverage clean-coverage bin