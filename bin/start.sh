#!/usr/bin/env bash
set -e

SECRET_PATH="/Users/uladpersonal/babajka-secret/secret-staging.json"

cd cli
go run *.go \
  --secretPath=$SECRET_PATH \
  --env="dev" \
  --enableSlack \
  --printReport
