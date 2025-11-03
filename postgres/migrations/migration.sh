#!/bin/bash
source .env

DB_HOST=${DB_HOST:-postgres-prod}
PG_PROD_PORT=${PG_PROD_PORT:-5432}
MIGRATION_DIR=${MIGRATION_DIR:-/root/migrations}

export MIGRATION_DSN="host=$DB_HOST port=$PG_PROD_PORT dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v