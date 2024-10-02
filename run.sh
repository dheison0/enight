#!/usr/bin/env bash

echo "Loading config..."
source .env
source api/.env


echo "Configuring database..."
echo "Running migrations..."
for f in $(find api/migrations -type f); do
  echo "Running for $f..."
  sqlite3 $DB_PATH "$(<$f)"
done

echo "Building web app..."
cd web
npm run build
cd ..

echo "Building API server..."
cd api
go build
cd ..
./api/api