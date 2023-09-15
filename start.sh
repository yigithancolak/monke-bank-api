#!/bin/bash

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$POSTGRES_URL" -verbose up

echo "starting the app"
exec "$@"