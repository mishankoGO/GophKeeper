package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewCards is a constructor for CardsServer interface instance.
func NewCards(repo interfaces.Repository, security *security.Security) *Cards {
	return &Cards{
		Repo:     repo,
		Security: *security,
	}
}

// Cards struct realizes CardsServer interface.
type Cards struct {
	pb.UnimplementedCardsServer
	Repo     interfaces.Repository // data storage
	Security security.Security     // cipher component
}

// Insert method encrypts and inserts data to db.
func (c *Cards) Insert(ctx context.Context, req *pb.InsertCardRequest) (*pb.InsertResponse, error) {
	// convert proto card to model card
	card, err := converters.PBCardToCard(req.User.UserId, req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto card to model card: %v", err)
	}

	// get card
	card_ := card.Card

	// create encoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	res := &pb.InsertResponse{IsInserted: false}

	// marshall into bytes
	err = encoder.Encode(string(card_))
	if err != nil {
		return res, status.Errorf(codes.Internal, "error encoding card: %v", err)
	}

	// encrypt data
	encData := c.Security.EncryptData(buf)

	// set encrypted card as Card
	card.Card = encData

	// insert new card to db
	err = c.Repo.InsertC(ctx, card)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting card: %v", err)
	}

	// set status
	res.IsInserted = true
	return res, nil
}

// Get method gets and decrypts data from db.
func (c *Cards) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetCardResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// get card from database
	card, err := c.Repo.GetC(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting card: %v", err)
	}

	// decrypt data
	decData, err := c.Security.DecryptData(card.Card)

	// set decrypted card to Card
	card.Card = bytes.Trim(decData, "\"\n")

	if card.Meta == nil {
		card.Meta = map[string]string{}
	}

	// convert card to proto card
	pbCard, err := converters.CardToPBCard(card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	// create response
	res := &pb.GetCardResponse{Card: pbCard}

	return res, nil
}

// Update method encrypts new binary file and updates record in db.
func (c *Cards) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	// convert proto user to model use
	user := converters.PBUserToUser(req.GetUser())

	// convert proto card to model card
	card, err := converters.PBCardToCard(user.UserID, req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	// get card
	card_ := card.Card

	// create encoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	// marshall into bytes
	err = encoder.Encode(string(card_))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error encoding card: %v", err)
	}

	// encrypt data
	encData := c.Security.EncryptData(buf)

	// set encrypted card as Card
	card.Card = encData

	// update record in db
	updatedCard, err := c.Repo.UpdateC(ctx, card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating card: %v", err)
	}

	// covert model card to proto card
	pbCard, err := converters.CardToPBCard(updatedCard)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	// create response
	res := &pb.UpdateCardResponse{Card: pbCard}

	return res, nil
}

// Delete method deletes binary file record from db.
func (c *Cards) Delete(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// create response
	res := &pb.DeleteResponse{Ok: false}

	// delete record
	err := c.Repo.DeleteC(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting card: %v", err)
	}

	// set status
	res.Ok = true
	return res, nil
}
