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

run-server:
	go run cmd/server/main.go

run-server-windows:
	./bin/server_windows

run-server-mac:
	./bin/server_mac

run-server-linux:
	./bin/server_linux

run-client-windows:
	./bin/client_windows

run-client-mac:
	./bin/client_mac

run-client-linux:
	./bin/client_linux

build-server:
	GOOS=windows go build -o ./bin/server_windows ./cmd/server/main.go && \
	GOOS=darwin GOARCH=amd64 go build -o ./bin/server_mac ./cmd/server/main.go && \
	GOOS=linux GOARCH=ppc64 go build -o ./bin/server_linux ./cmd/server/main.go

build-client:
	GOOS=windows go build -o ./bin/client_windows -v -ldflags="-X 'github.com/mishankoGO/GophKeeper/internal/cli/build_version.Version=v1.0.0' -X 'github.com/mishankoGO/GophKeeper/internal/cli/build_version.BuildDate=$(shell date)'" ./cmd/client/main.go && \
	GOOS=darwin GOARCH=amd64 go build -o ./bin/client_mac -v -ldflags="-X 'github.com/mishankoGO/GophKeeper/internal/cli/build_version.Version=v1.0.0' -X 'github.com/mishankoGO/GophKeeper/internal/cli/build_version.BuildDate=$(shell date)'"  ./cmd/client/main.go && \
	GOOS=linux GOARCH=ppc64 go build -o ./bin/client_linux -v -ldflags="-X 'github.com/mishankoGO/GophKeeper/internal/cli/build_version.Version=v1.0.0' -X 'github.com/mishankoGO/GophKeeper/internal/cli/build_version.BuildDate=$(shell date)'" ./cmd/client/main.go

evans:
	 evans -r repl -p 8080

docs:
	godoc -http=:8081

.PHONY: pg-run migrateup dropdb migrateup migratedown proto-gen run-server run-server-windows run-server-mac run-server-linux run-client-windows run-client-mac run-client-linux build-server build-client docs