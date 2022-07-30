POSTGRESQL_URL = 'postgres://postgres:537j04222@localhost:5432/postgres?sslmode=disable'


migrate_up:
	migrate -database $(POSTGRESQL_URL) -path db/migrations up

migrate_down:
	migrate -database $(POSTGRESQL_URL) -path db/migrations down
test:
	go test -v -cover ./...
sqlc:
	sqlc generate
server:
	go run main.go
mock:
	mockgen -destination transactions/db/mock/store.go  transactions/db/sqlc Store 