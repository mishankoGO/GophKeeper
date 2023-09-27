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

// Credentials struct realizes CredentialServer interface.
type Credentials struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedCredentialsServer
	Repo interfaces.Repository
}

func (c *Credentials) Register(ctx context.Context, req *pb.RegisterRequest) *pb.RegisterResponse {
	// hash password
	hashPass := hash.HashPass([]byte(req.Cred.Password))

	// convert pb credential to model
	cred := converters.PBCredentialToCredential(req.Cred)
	cred.Password = hashPass

	// register user
	u, err := c.Repo.Register(ctx, cred)
	if err != nil {
		log.Println(status.Errorf(codes.Internal, "error registering user: %v", err))
		return nil
	}

	// convert model user to pb user
	pbuser := converters.UserToPBUser(u)

	var res = &pb.RegisterResponse{User: pbuser}

	return res
}
