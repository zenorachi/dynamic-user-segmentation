#!/bin/bash

MAX_RETRIES=10
RETRY_INTERVAL=1

retries=0

until pg_isready -h "$DB_HOST" -U "$DB_USER" || [ "$retries" -eq "$MAX_RETRIES" ]; do
    echo "Waiting for the database to become available... Attempt $((retries + 1)) of $MAX_RETRIES"
    sleep "$RETRY_INTERVAL"
    ((retries++))
done

if [ "$retries" -eq "$MAX_RETRIES" ]; then
    echo "Failed to connect to the database after $MAX_RETRIES attempts. Exiting."
    exit 1
fi

echo "$(date '+%Y-%m-%d %H:%M:%S') - Database is now available."