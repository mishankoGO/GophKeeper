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

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
)

// CardsClient contains cards client service.
type CardsClient struct {
	service pb.CardsClient
	repo    interfaces.Repository
	offline bool
}

// NewCardsClient creates new cards client.
func NewCardsClient(cc *grpc.ClientConn, repo interfaces.Repository) *CardsClient {
	if cc != nil {
		service := pb.NewCardsClient(cc)
		return &CardsClient{service: service, repo: repo, offline: false}
	}
	return &CardsClient{repo: repo, offline: true}
}

// Insert method inserts new card.
func (c *CardsClient) Insert(ctx context.Context, req *pb.InsertCardRequest) (*pb.InsertResponse, error) {
	card, err := converters.PBCardToCard(req.GetUser().GetUserId(), req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto card to model card: %v", err)
	}
	err = c.repo.InsertC(card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting card: %v", err)
	}
	if !c.offline {
		resp, err := c.service.Insert(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error inserting card: %v", err)
		}

		return resp, nil
	}

	return &pb.InsertResponse{IsInserted: true}, nil
}

// Get method retrieves card information.
func (c *CardsClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetCardResponse, error) {
	card, err := c.repo.GetC(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving card: %v", err)
	}
	protoCard, err := converters.CardToPBCard(card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting model card to proto card: %v", err)
	}

	//resp, err := c.service.Get(ctx, req)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "error getting card information: %v", err)
	//}
	//
	//return resp, nil

	return &pb.GetCardResponse{Card: protoCard}, nil
}

// Update method updates card information.
func (c *CardsClient) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	mCard, err := converters.PBCardToCard(req.GetUser().GetUserId(), req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto card to model card: %v", err)
	}
	_, err = c.repo.UpdateC(mCard)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating card: %v", err)
	}
	if !c.offline {
		resp, err := c.service.Update(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error updating card information: %v", err)
		}

		return resp, nil
	}

	return &pb.UpdateCardResponse{Card: req.GetCard()}, nil
}

// Delete method deletes card.
func (c *CardsClient) Delete(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteResponse, error) {
	err := c.repo.DeleteC(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting binary file: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Delete(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error deleting card information: %v", err)
		}

		return resp, nil
	}
	return &pb.DeleteResponse{Ok: true}, nil
}

// List method to list all cards.
func (c *CardsClient) List() {

}
