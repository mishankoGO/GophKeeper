package handlers

import (
	"context"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
)

type Cards struct {
	pb.UnimplementedCardsServer
	Repo repository.Repository
}

func (c *Cards) Insert(ctx context.Context, req *pb.InsertCardRequest) (*pb.InsertResponse, error) {
	return nil, nil
}

func (c *Cards) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetCardResponse, error) {
	return nil, nil
}

func (c *Cards) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	return nil, nil
}
func (c *Cards) Delete(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteResponse, error) {
	return nil, nil
}
