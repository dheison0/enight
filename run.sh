#!/usr/bin/env bash

ROOT=$(dirname "$(realpath "$0")")
ROOT=$(realpath "$ROOT")

log() {
  echo -e "[`date +%H:%M:%S`] $@"
}

log "Building all system..."
cd "$ROOT/web"
npm run build &>/dev/null &
cd "$ROOT/server"
go build -ldflags="-extldflags=-w -s" || exit


log "Running server..."
mkdir "$ROOT/tmp"
export WEB_FILES="$ROOT/web/dist"
export DB_PATH="$ROOT/tmp/system.sqlite3"
export BOT_DB_PATH="$ROOT/tmp/bot.sqlite3"
export DEBUG=true
cd "$ROOT"
exec "$ROOT/server/server"
