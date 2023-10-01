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

// TextsClient contains text client service.
type TextsClient struct {
	service pb.TextsClient
}

// NewTextsClient creates new Texts client.
func NewTextsClient(cc *grpc.ClientConn) *TextsClient {
	service := pb.NewTextsClient(cc)
	return &TextsClient{service: service}
}

// Insert method inserts new Texts.
func (c *TextsClient) Insert(ctx context.Context, req *pb.InsertTextRequest) (*pb.InsertResponse, error) {
	resp, err := c.service.Insert(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting text: %v", err)
	}

	return resp, nil
}

// Get method retrieves text information.
func (c *TextsClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetTextResponse, error) {
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting text information: %v", err)
	}

	return resp, nil
}

// Update method updates text information.
func (c *TextsClient) Update(ctx context.Context, req *pb.UpdateTextRequest) (*pb.UpdateTextResponse, error) {
	resp, err := c.service.Update(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating text information: %v", err)
	}

	return resp, nil
}

// Delete method deletes text.
func (c *TextsClient) Delete(ctx context.Context, req *pb.DeleteTextRequest) (*pb.DeleteResponse, error) {
	resp, err := c.service.Delete(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting text information: %v", err)
	}

	return resp, nil
}
