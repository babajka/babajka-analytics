#!/usr/bin/env bash
set -e

bold=$(tput bold)
normal=$(tput sgr0)

MODE=$1

if [[ $MODE != "dev" && $MODE != "prod" ]]; then
  echo '[FAIL] Mode (dev or prod) must be provided.'
  exit 1
fi

if [[ $MODE == "prod" ]]; then
  echo "You're about to deploy to ${bold}prod${normal}. Are you sure? (put the number)"
  select yn in "Yes" "No"; do
    case $yn in
      "Yes" ) break;;
      "No" ) exit 0;;
    esac
  done
fi

# This locally builds binary which will run on remote server.
GOOS=linux GOARCH=amd64 go build -o build/analytics-linux ./cli
echo "[OK] Binary is built locally"

if [[ $MODE == "dev" ]]; then
  HOST="wir-dev@dev.wir.by"
elif [[ $MODE == "prod" ]]; then
  HOST="wir-prod@wir.by"
fi

BIN_FOLDER="/home/wir-$MODE/deployed/analytics"

ssh "${HOST}" "mkdir -p \"${BIN_FOLDER}\""
echo '[OK] Remote binary folder ensured to exist'

scp ./build/analytics-linux "${HOST}:${BIN_FOLDER}/babajka-analytics"
echo "[OK] Analytics binary pushed to ${MODE} server"
