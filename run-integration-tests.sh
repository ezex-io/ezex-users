#!/bin/bash
set -e

# load environment variables from .env.test
echo "Loading test environment from .env.test..."
if [ -f .env.test ]; then
  export $(cat .env.test | grep -v '^#' | xargs)
  echo "Environment loaded successfully"
else
  echo "Warning: .env.test file not found, using default values"
fi

# extract port from EZEX_USERS_DB_ADDRESS if POSTGRES_PORT is not set
if [ -z "$POSTGRES_PORT" ] && [ -n "$EZEX_USERS_DB_ADDRESS" ]; then
  export POSTGRES_PORT=$(echo $EZEX_USERS_DB_ADDRESS | cut -d ':' -f 2)
  echo "Extracted port $POSTGRES_PORT from EZEX_USERS_DB_ADDRESS"
fi

echo "Starting PostgreSQL for integration tests on port ${POSTGRES_PORT:-5433}..."
docker-compose -f docker-compose.test.yml up -d

echo "Waiting for PostgreSQL to be ready..."
sleep 5

echo "Running integration tests..."
go test -v ./internal/test/integration/...

echo "Cleaning up..."
docker-compose -f docker-compose.test.yml down

echo "Integration tests completed!" 