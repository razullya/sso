migrate create -ext sql -dir migrations -seq test 
go run ./cmd/sso/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./migrations

go run cmd/sso/main.go --config=./config/local.yaml  