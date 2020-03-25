#!/usr/bin/env bash
set -e

MODE=$1
if [ $MODE != "dev" ]; then
  echo '[FAIL] Mode (for now it is only DEV) must be provided explicitly'
  exit 1
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
