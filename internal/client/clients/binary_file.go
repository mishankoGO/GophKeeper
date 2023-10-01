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
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BinaryFilesClient contains binary files client service.
type BinaryFilesClient struct {
	service pb.BinaryFilesClient
}

// NewBinaryFilesClient creates new BinaryFiless client.
func NewBinaryFilesClient(cc *grpc.ClientConn) *BinaryFilesClient {
	service := pb.NewBinaryFilesClient(cc)
	return &BinaryFilesClient{service: service}
}

// Insert method inserts new BinaryFiles.
func (c *BinaryFilesClient) Insert(ctx context.Context, req *pb.InsertBinaryFileRequest) (*pb.InsertResponse, error) {
	resp, err := c.service.Insert(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
	}

	return resp, nil
}

// Get method retrieves binary file information.
func (c *BinaryFilesClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetBinaryFileResponse, error) {
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting binary file information: %v", err)
	}

	return resp, nil
}

// Update method updates binary file information.
func (c *BinaryFilesClient) Update(ctx context.Context, req *pb.UpdateBinaryFileRequest) (*pb.UpdateBinaryFileResponse, error) {
	resp, err := c.service.Update(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating binary file information: %v", err)
	}

	return resp, nil
}

// Delete method deletes binary file.
func (c *BinaryFilesClient) Delete(ctx context.Context, req *pb.DeleteBinaryFileRequest) (*pb.DeleteResponse, error) {
	resp, err := c.service.Delete(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting binary file information: %v", err)
	}

	return resp, nil
}
