package handlers

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"github.com/mishankoGO/GophKeeper/pkg/hash"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Users struct {
	pb.UnimplementedUsersServer
	Repo interfaces.Repository
}

func (u *Users) Login(ctx context.Context, req *pb.LoginRequest) *pb.LoginResponse {
	pbcred := req.Cred

	cred, user, err := u.Repo.Login(ctx, pbcred.Login)
	if err != nil {
		log.Println(status.Errorf(codes.Internal, "error getting user id: %v", err))
		return nil
	}

	// hash password
	hashPass := hash.HashPass([]byte(pbcred.Password))

	if cred.Password != hashPass {
		log.Println(status.Error(codes.PermissionDenied, "invalid credentials"))
		return nil
	}
	log.Println("Logged in successfully!")

	pbuser := converters.UserToPBUser(user)
	var res = &pb.LoginResponse{User: pbuser}
	return res
}
