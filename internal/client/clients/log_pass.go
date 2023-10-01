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

// LogPassesClient contains log pass client service.
type LogPassesClient struct {
	service pb.LogPassesClient
}

// NewLogPassesClient creates new LogPasses client.
func NewLogPassesClient(cc *grpc.ClientConn) *LogPassesClient {
	service := pb.NewLogPassesClient(cc)
	return &LogPassesClient{service: service}
}

// Insert method inserts new LogPasses.
func (c *LogPassesClient) Insert(ctx context.Context, req *pb.InsertLogPassRequest) (*pb.InsertResponse, error) {
	resp, err := c.service.Insert(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting log pass: %v", err)
	}

	return resp, nil
}

// Get method retrieves log pass information.
func (c *LogPassesClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetLogPassResponse, error) {
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting log pass information: %v", err)
	}

	return resp, nil
}

// Update method updates log pass information.
func (c *LogPassesClient) Update(ctx context.Context, req *pb.UpdateLogPassRequest) (*pb.UpdateLogPassResponse, error) {
	resp, err := c.service.Update(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating log pass information: %v", err)
	}

	return resp, nil
}

// Delete method deletes log pass.
func (c *LogPassesClient) Delete(ctx context.Context, req *pb.DeleteLogPassRequest) (*pb.DeleteResponse, error) {
	resp, err := c.service.Delete(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting log pass information: %v", err)
	}

	return resp, nil
}
