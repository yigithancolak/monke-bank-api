postgresrun:
	docker run --rm --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=monke_bank -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres monke_bank

adminer:
	@sh -c 'docker run --rm --link postgres12:db -p 8080:8080 adminer'

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/monke_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/monke_bank?sslmode=disable" -verbose down

test:
	go test -v -cover -short ./...

proto:
	rm -f pb/*.go && \
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto


.PHONY: postgresrun createdb dropdb migrateup migratedown adminer,test, proto