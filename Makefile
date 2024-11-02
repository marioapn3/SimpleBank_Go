postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 4321:5432 -d postgres:17.0	

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:4321/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:4321/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/marioapn3/simplebank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test