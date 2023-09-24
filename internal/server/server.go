package server

import (
	"context"
	// импортируем пакет со сгенерированными protobuf-файлами
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
)

type Server interface {
	Credentials
}

type Credentials interface {
	Register(ctx context.Context, pbcred *pb.RegisterRequest) (*pb.User, error)
}
