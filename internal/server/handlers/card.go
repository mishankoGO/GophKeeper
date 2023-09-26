package handlers

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Cards struct {
	pb.UnimplementedCardsServer
	Repo repository.Repository
}

func (c *Cards) Insert(ctx context.Context, req *pb.InsertCardRequest) (*pb.InsertResponse, error) {
	card, err := converters.PBCardToCard(req.User.UserId, req.GetCard())
	res := &pb.InsertResponse{IsInserted: false}
	if err != nil {
		return res, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	err = c.Repo.InsertC(ctx, card)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting card: %v", err)
	}

	res.IsInserted = true
	return res, nil
}
func (c *Cards) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetCardResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	card, err := c.Repo.GetC(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting card: %v", err)
	}

	pbCard, err := converters.CardToPBCard(card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	res := &pb.GetCardResponse{Card: pbCard}

	return res, nil
}
func (c *Cards) Update(ctx context.Context, req *pb.UpdateCardRequest) (*pb.UpdateCardResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	card, err := converters.PBCardToCard(user.UserID, req.GetCard())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	card_, err := c.Repo.UpdateC(ctx, card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating card: %v", err)
	}

	pbCard, err := converters.CardToPBCard(card_)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting card: %v", err)
	}

	res := &pb.UpdateCardResponse{Card: pbCard}

	return res, nil
}
func (c *Cards) Delete(ctx context.Context, req *pb.DeleteCardRequest) (*pb.DeleteResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	res := &pb.DeleteResponse{Ok: false}

	err := c.Repo.DeleteC(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting card: %v", err)
	}

	res.Ok = true
	return res, nil
}
