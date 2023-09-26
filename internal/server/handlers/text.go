package handlers

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Texts struct {
	pb.UnimplementedTextsServer
	Repo repository.Repository
}

func (t *Texts) Insert(ctx context.Context, req *pb.InsertTextRequest) (*pb.InsertResponse, error) {
	text, err := converters.PBTextToText(req.User.UserId, req.Text)
	res := &pb.InsertResponse{IsInserted: false}
	if err != nil {
		return res, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	err = t.Repo.InsertT(ctx, text)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting text: %v", err)
	}

	res.IsInserted = true
	return res, nil
}

func (t *Texts) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetTextResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	text, err := t.Repo.GetT(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting text: %v", err)
	}

	pbText, err := converters.TextToPBText(text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	res := &pb.GetTextResponse{Text: pbText}

	return res, nil
}

func (t *Texts) Update(ctx context.Context, req *pb.UpdateTextRequest) (*pb.UpdateTextResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	text, err := converters.PBTextToText(user.UserID, req.GetText())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	text_, err := t.Repo.UpdateT(ctx, text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating text: %v", err)
	}

	pbText, err := converters.TextToPBText(text_)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	res := &pb.UpdateTextResponse{Text: pbText}

	return res, nil
}

func (t *Texts) Delete(ctx context.Context, req *pb.DeleteTextRequest) (*pb.DeleteResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	res := &pb.DeleteResponse{Ok: false}

	err := t.Repo.DeleteT(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting text: %v", err)
	}

	res.Ok = true
	return res, nil
}
