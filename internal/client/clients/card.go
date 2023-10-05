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
	"bytes"
	"context"
	"encoding/json"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
)

// CardsClient contains cards client service.
type CardsClient struct {
	service  pb.CardsClient
	Security *security.Security
	repo     interfaces.Repository
	offline  bool
}

// NewCardsClient creates new cards client.
func NewCardsClient(cc *grpc.ClientConn, repo interfaces.Repository, security *security.Security) *CardsClient {
	if cc != nil {
		service := pb.NewCardsClient(cc)
		return &CardsClient{service: service, repo: repo, Security: security, offline: false}
	}
	return &CardsClient{repo: repo, Security: security, offline: true}
}

// Insert method inserts new card.
func (c *CardsClient) Insert(ctx context.Context, req *pb.InsertCardRequest) (*pb.InsertResponse, error) {
	card, err := converters.PBCardToCard(req.GetUser().GetUserId(), req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto card to model card: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(string(card.Card))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error encoding card: %v", err)
	}
	encData := c.Security.EncryptData(buf)

	// set encrypted card as Card
	card.Card = encData

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
	if c.offline {
		card, err := c.repo.GetC(req.GetName())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error retrieving card: %v", err)
		}

		// decrypt data
		decData, err := c.Security.DecryptData(card.Card)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
		}

		// set decrypted card to Card
		card.Card = bytes.Trim(decData, "\"\n")

		protoCard, err := converters.CardToPBCard(card)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error converting model card to proto card: %v", err)
		}
		return &pb.GetCardResponse{Card: protoCard}, nil
	}
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting card information: %v", err)
	}
	return resp, nil
}

// Update method updates card information.
func (c *CardsClient) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	mCard, err := converters.PBCardToCard(req.GetUser().GetUserId(), req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto card to model card: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(string(mCard.Card))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error encoding card: %v", err)
	}
	encData := c.Security.EncryptData(buf)

	// set encrypted card as Card
	mCard.Card = encData

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
