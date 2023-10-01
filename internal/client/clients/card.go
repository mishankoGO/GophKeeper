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

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
)

// CardsClient contains cards client service.
type CardsClient struct {
	service pb.CardsClient
}

// NewCardsClient creates new cards client.
func NewCardsClient(cc *grpc.ClientConn) *CardsClient {
	service := pb.NewCardsClient(cc)
	return &CardsClient{service: service}
}

// Insert method inserts new card.
func (c *CardsClient) Insert(ctx context.Context, req *pb.InsertCardRequest) (*pb.InsertResponse, error) {
	resp, err := c.service.Insert(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting card: %v", err)
	}

	return resp, nil
}

// Get method retrieves card information.
func (c *CardsClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetCardResponse, error) {
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting card information: %v", err)
	}

	return resp, nil
}

// Update method updates card information.
func (c *CardsClient) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	resp, err := c.service.Update(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating card information: %v", err)
	}

	return resp, nil
}

// Delete method deletes card.
func (c *CardsClient) Delete(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteResponse, error) {
	resp, err := c.service.Delete(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting card information: %v", err)
	}

	return resp, nil
}
