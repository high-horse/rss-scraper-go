.PHONY: goose

run :
	@go run .

sqlc_i: 
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

goose_i : 
	go install github.com/pressly/goose/v3/cmd/goose@latest

sqlc :	
	sqlc generate

goose_up:
	cd sql/schema && goose postgres postgres://postgres:root@localhost:5432/go-rss up

goose_down:
	cd sql/schema && goose postgres postgres://postgres:root@localhost:5432/go-rss down
