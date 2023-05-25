installpostgres:
	docker pull postgres:12-alpine

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root eightcig

dropdb:
	docker exec -it postgres12 dropdb eightcig

migrateup:
	./utils/migrate -path db/migration -database "postgresql://root:secret@localhost:5432/eightcig?sslmode=disable" --verbose up

migratedown:
	./utils/migrate -path db/migration -database "postgresql://root:secret@localhost:5432/eightcig?sslmode=disable" --verbose up

runapi:
	go run ./cmd/api/

.PHONY: postgres createdb dropdb migrateup migratedown installpostgres
