#!/usr/bin/env bash

ROOT=$(dirname "$(realpath "$0")")
ROOT=$(realpath "$ROOT")

log() {
  echo -e "[`date +%H:%M:%S`] $@"
}

log "Building all system..."
if [ $(nproc) -gt 2 ]; then
  cd "$ROOT/web"
  npm run build &>/dev/null &
fi
cd "$ROOT/server"
CGO_ENABLED=1 go build -ldflags="-extldflags=-w -s" || exit


log "Running server..."
mkdir "$ROOT/tmp"
export WEB_FILES_PATH="$ROOT/web/dist"
export DB_PATH="$ROOT/tmp/system.sqlite3"
export BOT_DB_PATH="$ROOT/tmp/bot.sqlite3"
export DEBUG=true
cd "$ROOT"
exec "$ROOT/server/server"
