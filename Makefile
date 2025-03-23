APP_NAME=go-auth-service
BUILD_DIR=bin

.PHONY: build run clean

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go

run: build
	./$(BUILD_DIR)/$(APP_NAME)

migrate:
	go run internal/migration/migration.go

clean:
	rm -rf $(BUILD_DIR)
