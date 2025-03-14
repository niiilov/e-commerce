#!/bin/bash

APP_SRC="./cmd/main.go"

if [ ! -f "$APP_SRC" ]; then
  echo "Error: source file $APP_SRC not exists!"
  exit 1
fi

echo "Compiling..."
go run "$APP_SRC"
