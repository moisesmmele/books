ifneq (,$(wildcard .env))
	include .env
	export
endif

DSN="host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable timezone=utc connect_timeout=5"
BINARY_NAME=books-api

build:
	@echo "building the backend..."
	go build -i ${BINARY_NAME} ./cmd/api/
	@echo "built."

run: build
	@echo "starting the backend..."
	@env DSN=${DSN} ./${BINARY_NAME} &
	@echo "started."

clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "cleaned."

start: run

stop:
	@echo "stopping the backend..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "stopped."