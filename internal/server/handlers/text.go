package handlers

import (
	"context"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
)

type Texts struct {
	pb.UnimplementedTextsServer
	Repo repository.Repository
}

func (lp *Texts) Insert(ctx context.Context, req *pb.InsertTextRequest) (*pb.InsertResponse, error) {
	return nil, nil
}

func (lp *Texts) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetTextResponse, error) {
	return nil, nil
}

func (lp *Texts) Update(ctx context.Context, req *pb.UpdateTextRequest) (*pb.UpdateTextResponse, error) {
	return nil, nil
}
func (lp *Texts) Delete(ctx context.Context, req *pb.DeleteTextRequest) (*pb.DeleteResponse, error) {
	return nil, nil
}
