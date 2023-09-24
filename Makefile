pg-run:
	docker run --name gophkeeper-pg -p 5432:5432 -e POSTGRES_USER=gophkeeperuser -e POSTGRES_PASSWORD=gophkeeperpwd -e POSTGRES_DB=gopgkeeperdb -d postgres

dropdb:
	docker exec -it gophkeeper-pg dropdb gophkeeperdb

migrateup:
	migrate -path migrations -database "postgresql://gophkeeperuser:gophkeeperpwd@localhost:5432/gophkeeperdb?sslmode=disable" -verbose up 1

migratedown:
	migrate -path migrations -database "postgresql://gophkeeperuser:gophkeeperpwd@localhost:5432/gophkeeperdb?sslmode=disable" -verbose down 1

proto-gen:
	\protoc -I=api/  --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative api/binary_file.proto && \
    \protoc -I=api/  --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative api/card.proto && \
    \protoc -I=api/  --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative api/log_pass.proto && \
    \protoc -I=api/  --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative api/text.proto && \
    \protoc -I=api/  --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative api/user.proto && \
    \protoc -I=api/  --go_out=internal/grpc --go_opt=paths=source_relative --go-grpc_out=internal/grpc --go-grpc_opt=paths=source_relative api/registration.proto

.PHONY: pg-run migrateup dropdb migrateup migratedown