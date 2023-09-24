package handlers

import (
	"context"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
)

type BinaryFiles struct {
	pb.UnimplementedBinaryFilesServer
	Repo repository.Repository
}

func (bf *BinaryFiles) Insert(ctx context.Context, req *pb.InsertBinaryFileRequest) (*pb.InsertResponse, error) {
	return nil, nil
}

func (bf *BinaryFiles) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetBinaryFileResponse, error) {
	return nil, nil
}

func (bf *BinaryFiles) Update(ctx context.Context, req *pb.UpdateBinaryFileRequest) (*pb.UpdateBinaryFileResponse, error) {
	return nil, nil
}
func (bf *BinaryFiles) Delete(ctx context.Context, req *pb.DeleteBinaryFileRequest) (*pb.DeleteResponse, error) {
	return nil, nil
}
