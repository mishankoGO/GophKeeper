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
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/pkg/util"
)

// CardsClient contains cards client service.
type CardsClient struct {
	service  pb.CardsClient
	Security *security.Security
	repo     interfaces.Repository
	offline  bool
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

	// encrypt data
	var buf bytes.Buffer
	buf.Write(card.Card)

	encData := c.Security.EncryptData(buf)

	// set encrypted card as Card
	card.Card = encData

	err = c.repo.InsertC(card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting card: %v", err)
	}

	if !c.offline {
		req.Card.Card = encData
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

	// decrypt data
	cc := resp.Card.Card
	decData, err := c.Security.DecryptData(cc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
	}

	// set decrypted card to card
	resp.Card.Card = bytes.Trim(decData, "\"\n")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving card: %v", err)
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
	buf.Write(mCard.Card)

	encData := c.Security.EncryptData(buf)

	// set encrypted card as Card
	mCard.Card = encData

	_, err = c.repo.UpdateC(mCard)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating card: %v", err)
	}

	if !c.offline {
		req.Card.Card = encData
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
func (c *CardsClient) List(ctx context.Context, req *pb.ListCardRequest) (*pb.ListCardResponse, []*cards.Cards, error) {
	cards, err := c.repo.ListC()
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "error listing cards: %v", err)
	}

	resp, err := c.service.List(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("error listing data from server: %v", err)
	}
	return resp, cards, err
}

// Sync method to sync cards between dbs.
func (c *CardsClient) Sync(ctx context.Context, req *pb.ListCardRequest) error {
	serverCs, clientCs, err := c.List(ctx, req)
	if err != nil {
		return err
	}

	// arrays of names for future syncing
	clientNames := make([]string, len(clientCs))
	serverNames := make([]string, len(serverCs.GetCards()))

	// flag which shows which db has the latest data.
	// if flag set to "server", it means server has fresher data.
	var dataPrimary string
	if c.offline {
		dataPrimary = "client"
	} else {
		dataPrimary = "server"
	}

	// update cycle
	for _, cc := range clientCs {
		cname := cc.Name
		clientNames = append(clientNames, cname)
		for _, sc := range serverCs.GetCards() {
			sname := sc.GetName()
			serverNames = append(serverNames, sname)

			// update common cards
			if sname == cname {
				err = c.updateCommonFiles(ctx, req, cc, sc)
				if err != nil {
					return err
				}
			}
		}
	}

	if dataPrimary == "server" {
		// insert missing server cards to client
		err = c.insertServerToClient(req, serverCs, clientNames)
		if err != nil {
			return err
		}

		// delete cards
		err = c.deleteFromClient(clientCs, serverNames)
		if err != nil {
			return err
		}
	} else if dataPrimary == "client" {
		// insert missing client cards to cards
		err = c.insertClientToServer(ctx, clientCs, serverNames, req)
		if err != nil {
			return err
		}

		// delete cards
		err = c.deleteFromServer(ctx, req, serverCs, clientNames)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetSecurity method to set security attribute.
func (c *CardsClient) SetSecurity(security *security.Security) {
	c.Security = security
}

func (c *CardsClient) updateCommonFiles(
	ctx context.Context,
	req *pb.ListCardRequest,
	cc *cards.Cards,
	sc *pb.Card) error {
	if cc.UpdatedAt.After(sc.UpdatedAt.AsTime()) {
		// convert card to proto card
		protoC, err := converters.CardToPBCard(cc)
		if err != nil {
			return err
		}

		// update server card
		reqS := &pb.UpdateCardRequest{User: req.GetUser(), Card: protoC}
		_, err = c.service.Update(ctx, reqS)
		if err != nil {
			return err
		}
	} else {
		// convert proto card to card
		card, err := converters.PBCardToCard(req.GetUser().GetUserId(), sc)
		if err != nil {
			return err
		}

		// update client card
		_, err = c.repo.UpdateC(card)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CardsClient) insertServerToClient(
	req *pb.ListCardRequest,
	serverCs *pb.ListCardResponse,
	clientNames []string) error {
	for _, sc := range serverCs.GetCards() {
		// convert proto card to model card
		card, err := converters.PBCardToCard(req.GetUser().GetUserId(), sc)
		if err != nil {
			return err
		}

		if !util.StringInSlice(sc.Name, clientNames) {
			// insert missing card to client db
			err = c.repo.InsertC(card)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CardsClient) deleteFromClient(clientCs []*cards.Cards, serverNames []string) error {
	for _, cc := range clientCs {
		if !util.StringInSlice(cc.Name, serverNames) {
			// delete cards absent in server
			err := c.repo.DeleteC(cc.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CardsClient) insertClientToServer(
	ctx context.Context,
	clientCs []*cards.Cards,
	serverNames []string,
	req *pb.ListCardRequest) error {
	for _, cc := range clientCs {
		// convert model card to proto card
		protoC, err := converters.CardToPBCard(cc)
		if err != nil {
			return err
		}

		if !util.StringInSlice(cc.Name, serverNames) {
			// insert missing card to server db
			reqS := &pb.InsertCardRequest{User: req.GetUser(), Card: protoC}
			_, err = c.service.Insert(ctx, reqS)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CardsClient) deleteFromServer(
	ctx context.Context,
	req *pb.ListCardRequest,
	serverCs *pb.ListCardResponse,
	clientNames []string) error {
	for _, sc := range serverCs.GetCards() {
		if !util.StringInSlice(sc.GetName(), clientNames) {
			// delete cards absent in client
			resD := &pb.DeleteCardRequest{User: req.GetUser(), Name: sc.GetName()}
			_, err := c.service.Delete(ctx, resD)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
