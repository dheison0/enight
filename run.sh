#!/usr/bin/env bash


function log() {
  echo -e "[`date +%H:%M:%S`] $@"
}

[[ -e '.env' ]] && source .env
[[ -e 'api/.env' ]] && source api/.env

log "Running migrations..."
for f in $(find api/migrations -type f); do
  log "\t $f..."
  sqlite3 "$DB_PATH" < "$f" || exit
done

log "Building all system..."
cd web
npm run build &>/dev/null &
cd ../api
go build -ldflags="-extldflags=-w -s" || exit
cd ..
log "Running server..."
./api/api