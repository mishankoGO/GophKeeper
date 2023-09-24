package handlers

import (
	"context"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
)

type LogPasses struct {
	pb.UnimplementedLogPassesServer
	Repo repository.Repository
}

func (lp *LogPasses) Insert(ctx context.Context, req *pb.InsertLogPassRequest) (*pb.InsertResponse, error) {
	return nil, nil
}

func (lp *LogPasses) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetLogPassResponse, error) {
	return nil, nil
}

func (lp *LogPasses) Update(ctx context.Context, req *pb.UpdateLogPassRequest) (*pb.UpdateLogPassResponse, error) {
	return nil, nil
}
func (lp *LogPasses) Delete(ctx context.Context, req *pb.DeleteLogPassRequest) (*pb.DeleteResponse, error) {
	return nil, nil
}
