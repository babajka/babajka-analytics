#!/usr/bin/env bash
set -e

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

PROJECT_NAME="babajka-analytics"
SRC_REMOTE_PATH="/home/wir-$MODE/go/src/github.com/babajka/$PROJECT_NAME"
HOST="wir-$MODE@$MODE.wir.by"
BIN_LOCATION="/home/wir-$MODE/deployed/analytics"

ssh $HOST "mkdir -p \"${SRC_REMOTE_PATH}\""
echo '[OK] src folder ensured to exist'

rsync -r --delete-after --exclude=.git . "$HOST:${SRC_REMOTE_PATH}/"
echo '[OK] Analytics sources pushed to server'

# ? do I need to explicitly install dependencies

ssh -l "wir-$MODE" $HOST 'bash -s' < bin/build-remote.sh $MODE
echo '[OK] Binary is built'
