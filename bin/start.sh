# SECRET_PATH=""
SECRET_PATH="/Users/uladbohdan/go/src/github.com/babajka/babajka-analytics/secret.json"

cd cli
go run *.go \
  --secretPath=$SECRET_PATH \
  --env="dev" \
  --enableSlack
