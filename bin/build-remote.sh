#!/usr/bin/env bash
set -e

MODE=$1

export GOPATH="/home/wir-$MODE/go"
echo "GOPATH ensured: $GOPATH"

cd "/home/wir-$MODE/go/src/github.com/babajka/babajka-analytics/cli"
/usr/local/go/bin/go build -o "/home/wir-$MODE/deployed/analytics/babajka-analytics"
