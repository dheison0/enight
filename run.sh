#!/usr/bin/env bash

ROOT=$(dirname "$(realpath "$0")")
cd "$ROOT"

log() {
  echo -e "[`date +%H:%M:%S`] $@"
}

if [[ -f '.env' ]]; then
  log "Loading .env file..."
  source .env
fi

if [[ -f 'server/.env' ]]; then
  log "Loading server/.env file..."
  source server/.env
fi

log "Running migrations..."
for file in server/database/migrations/*.sql; do
  log "\t $file..."
  sqlite3 "$DB_PATH" < "$file" || exit
done

log "Building all system..."
cd "$ROOT/web"
npm run build &>/dev/null &
cd "$ROOT/server"
go build -ldflags="-extldflags=-w -s" || exit
cd "$ROOT"

log "Running server..."
export WEB_FILES="${WEB_FILES:-$ROOT/web/dist}"
export DB_PATH="${DB_PATH:-$ROOT/db.sqlite3}"
export BOT_DB_PATH=${BOT_DB_PATH:-$ROOT/bot_db.sqlite3}
export DEBUG=${DEBUG:-true}
./server/server
