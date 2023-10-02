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

// TextsClient contains text client service.
type TextsClient struct {
	service pb.TextsClient
	repo    interfaces.Repository
	offline bool
}

// NewTextsClient creates new Texts client.
func NewTextsClient(cc *grpc.ClientConn, repo interfaces.Repository) *TextsClient {
	if cc != nil {
		service := pb.NewTextsClient(cc)
		return &TextsClient{service: service, repo: repo, offline: false}
	}
	return &TextsClient{repo: repo, offline: true}
}

// Insert method inserts new Texts.
func (c *TextsClient) Insert(ctx context.Context, req *pb.InsertTextRequest) (*pb.InsertResponse, error) {
	t, err := converters.PBTextToText(req.GetUser().GetUserId(), req.GetText())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto text to text: %v", err)
	}
	err = c.repo.InsertT(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting text: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Insert(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error inserting text: %v", err)
		}

		return resp, nil
	}
	return &pb.InsertResponse{IsInserted: true}, nil
}

// Get method retrieves text information.
func (c *TextsClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetTextResponse, error) {
	t, err := c.repo.GetT(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving text: %v", err)
	}
	protoT, err := converters.TextToPBText(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting model text to proto text: %v", err)
	}

	//if !c.offline {
	//	resp, err := c.service.Get(ctx, req)
	//	if err != nil {
	//		return nil, status.Errorf(codes.Internal, "error getting text information: %v", err)
	//	}
	//
	//	return resp, nil
	//}

	return &pb.GetTextResponse{Text: protoT}, nil
}

// Update method updates text information.
func (c *TextsClient) Update(ctx context.Context, req *pb.UpdateTextRequest) (*pb.UpdateTextResponse, error) {
	mText, err := converters.PBTextToText(req.GetUser().GetUserId(), req.GetText())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto text to model text: %v", err)
	}
	_, err = c.repo.UpdateT(mText)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating text: %v", err)
	}
	if !c.offline {
		resp, err := c.service.Update(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error updating text information: %v", err)
		}

		return resp, nil
	}
	return &pb.UpdateTextResponse{Text: req.GetText()}, nil
}

// Delete method deletes text.
func (c *TextsClient) Delete(ctx context.Context, req *pb.DeleteTextRequest) (*pb.DeleteResponse, error) {
	err := c.repo.DeleteT(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting text: %v", err)
	}
	if !c.offline {
		resp, err := c.service.Delete(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error deleting text information: %v", err)
		}

		return resp, nil
	}
	return &pb.DeleteResponse{Ok: true}, nil
}
