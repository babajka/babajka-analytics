SECRET_PATH=""

cd cli
go run *.go \
  --secretPath=$SECRET_PATH \
  --env="dev" \
  --enableSlack
