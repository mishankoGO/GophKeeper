package handlers

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogPasses struct {
	pb.UnimplementedLogPassesServer
	Repo repository.Repository
}

func (lp *LogPasses) Insert(ctx context.Context, req *pb.InsertLogPassRequest) (*pb.InsertResponse, error) {
	logPass, err := converters.PBLogPassToLogPass(req.User.UserId, req.GetLogPass())
	res := &pb.InsertResponse{IsInserted: false}
	if err != nil {
		return res, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	err = lp.Repo.InsertLP(ctx, logPass)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting log pass: %v", err)
	}

	res.IsInserted = true
	return res, nil
}

func (lp *LogPasses) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetLogPassResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	logPass, err := lp.Repo.GetLP(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting log pass: %v", err)
	}

	pbLogPass, err := converters.LogPassToPBLogPass(logPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	res := &pb.GetLogPassResponse{LogPass: pbLogPass}

	return res, nil
}

func (lp *LogPasses) Update(ctx context.Context, req *pb.UpdateLogPassRequest) (*pb.UpdateLogPassResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	logPass, err := converters.PBLogPassToLogPass(user.UserID, req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	logPass_, err := lp.Repo.UpdateLP(ctx, logPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating log pass: %v", err)
	}

	pbLogPass, err := converters.LogPassToPBLogPass(logPass_)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	res := &pb.UpdateLogPassResponse{LogPass: pbLogPass}

	return res, nil
}

func (lp *LogPasses) Delete(ctx context.Context, req *pb.DeleteLogPassRequest) (*pb.DeleteResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	res := &pb.DeleteResponse{Ok: false}

	err := lp.Repo.DeleteLP(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting log pass: %v", err)
	}

	res.Ok = true
	return res, nil
}
