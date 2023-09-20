pg-run:
	docker run --name gophkeeper-pg -p 5432:5432 -e POSTGRES_USER=gophkeeperuser -e POSTGRES_PASSWORD=gophkeeperpwd -e POSTGRES_DB=gopgkeeperdb -d postgres

dropdb:
	docker exec -it gophkeeper-pg dropdb gophkeeperdb

migrateup:
	migrate -path migrations -database "postgresql://gophkeeperuser:gophkeeperpwd@localhost:5432/gophkeeperdb?sslmode=disable" -verbose up 1

migratedown:
	migrate -path migrations -database "postgresql://gophkeeperuser:gophkeeperpwd@localhost:5432/gophkeeperdb?sslmode=disable" -verbose down 1

.PHONY: pg-run migrateup dropdb migrateup migratedown