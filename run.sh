#!/usr/bin/env bash

ROOT=$(dirname "$(realpath "$0")")
ROOT=$(realpath "$ROOT")
TMP_DIR="$ROOT/tmp"

function log() {
  echo -e "[`date +%H:%M:%S`] $@"
}

mkdir "$TMP_DIR" &>/dev/null

if [ $(nproc) -gt 2 ] || [ "$1" == "build-web" ]; then
  log "Building web..."
  cd "$ROOT/web"
  time npm run build -- --emptyOutDir --outDir "$TMP_DIR/www"
fi
log "Building server..."
cd "$ROOT/server"
CGO_ENABLED=0 time go build -ldflags="-w -s" -v -o "$TMP_DIR/server" || exit

log "Running server..."
export DEBUG=true
cd "$TMP_DIR"
exec "./server"
