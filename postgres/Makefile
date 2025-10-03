include .env

LOCAL_BIN := $(CURDIR)/bin

LOCAL_MIGRATION_DIR = $(MIGRATION_DIR)
LOCAL_MIGRATION_DSN = "host=$(HOST) port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

ifeq ($(OS),Windows_NT)
	# Windows
	SET_GOBIN = set GOBIN=$(LOCAL_BIN)&&
else
	# Linux/Mac
	SET_GOBIN = GOBIN=$(LOCAL_BIN)
endif

install-deps:
	$(SET_GOBIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

local-migration-status:
	./bin/goose.exe -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) status -v

local-migration-up:
	./bin/goose.exe -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

local-migration-down:
	./bin/goose.exe -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v
