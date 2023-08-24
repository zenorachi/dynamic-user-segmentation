#!/bin/bash

# Wait for the db-connection to become available

until pg_isready -h "$DB_HOST" -U "$DB_USER"; do
    echo "Waiting for the database to become available..."
    sleep 1
done

echo "DB started"