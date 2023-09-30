package handlers

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"github.com/mishankoGO/GophKeeper/pkg/hash"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewCredentials(repo interfaces.Repository) *Credentials {
	return &Credentials{Repo: repo}
}

// Credentials struct realizes CredentialServer interface.
type Credentials struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedCredentialsServer
	Repo interfaces.Repository
}

func (c *Credentials) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// hash password
	hashPass := hash.HashPass([]byte(req.Cred.Password))

	// convert pb credential to model
	cred := converters.PBCredentialToCredential(req.Cred)

	// set password
	cred.Password = hashPass

	// register user
	u, err := c.Repo.Register(ctx, cred)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error registering user: %v", err)
	}

	// convert model user to pb user
	pbUser := converters.UserToPBUser(u)

	// create response
	var res = &pb.RegisterResponse{User: pbUser}

	return res, nil
}
