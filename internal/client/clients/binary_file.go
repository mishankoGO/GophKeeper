// Package clients contains clients for all the operations.
// It has:
// CredentialsClient for registration.
// UsersClient for login in.
// CardsClient to operate with bank cards.
// LogPassesClient to operate with log passes.
// TextsClient to operate with texts.
// BinaryFilesClient to operate with binary files.
package clients

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BinaryFilesClient contains binary files client service.
type BinaryFilesClient struct {
	service pb.BinaryFilesClient
	repo    interfaces.Repository
	offline bool
}

// NewBinaryFilesClient creates new BinaryFiles client.
func NewBinaryFilesClient(cc *grpc.ClientConn, repo interfaces.Repository) *BinaryFilesClient {
	if cc != nil {
		service := pb.NewBinaryFilesClient(cc)
		return &BinaryFilesClient{service: service, repo: repo, offline: false}

	}
	return &BinaryFilesClient{repo: repo, offline: true}
}

// Insert method inserts new BinaryFiles.
func (c *BinaryFilesClient) Insert(ctx context.Context, req *pb.InsertBinaryFileRequest) (*pb.InsertResponse, error) {
	bf, err := converters.PBBinaryFileToBinaryFile(req.GetUser().GetUserId(), req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto binary file to binary file: %w", err)
	}
	err = c.repo.InsertBF(bf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Insert(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
		}
		return resp, nil
	}

	return &pb.InsertResponse{IsInserted: true}, nil

}

// Get method retrieves binary file information.
func (c *BinaryFilesClient) Get(req *pb.GetRequest) (*pb.GetBinaryFileResponse, error) {
	bf, err := c.repo.GetBF(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving binary file: %v", err)
	}
	protoBF, err := converters.BinaryFileToPBBinaryFile(bf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file to proto binary file: %v", err)
	}

	//if !c.offline {
	//	resp, err := c.service.Get(ctx, req)
	//	if err != nil {
	//		return nil, status.Errorf(codes.Internal, "error getting binary file information: %v", err)
	//	}
	//	return resp, nil
	//}

	return &pb.GetBinaryFileResponse{File: protoBF}, nil

}

// Update method updates binary file information.
func (c *BinaryFilesClient) Update(ctx context.Context, req *pb.UpdateBinaryFileRequest) (*pb.UpdateBinaryFileResponse, error) {
	mFile, err := converters.PBBinaryFileToBinaryFile(req.GetUser().GetUserId(), req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto binary file to model binary file: %v", err)
	}
	_, err = c.repo.UpdateBF(mFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating binary file: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Update(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error updating binary file information: %v", err)
		}
		return resp, nil
	}

	return &pb.UpdateBinaryFileResponse{File: req.GetFile()}, nil
}

// Delete method deletes binary file.
func (c *BinaryFilesClient) Delete(ctx context.Context, req *pb.DeleteBinaryFileRequest) (*pb.DeleteResponse, error) {
	err := c.repo.DeleteBF(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting binary file: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Delete(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error deleting binary file information: %v", err)
		}

		return resp, nil
	}

	return &pb.DeleteResponse{Ok: true}, nil
}
