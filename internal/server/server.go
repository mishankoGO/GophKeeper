package server

import (
	// импортируем пакет со сгенерированными protobuf-файлами
	pb "github.com/mishankoGO/GophKeeper/api"
	"github.com/mishankoGO/GophKeeper/internal/grpc/registration"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Credentials struct realizes CredentialServer interface.
type Credentials struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedCredentialsServer
	Repo repository.Repository
}

func (c *Credentials) Register(credential *api.Credential) (*api.User, error) {
	u, err := c.Repo.Register(credential)
	if err != nil {
		return u, status.Error(codes.Internal, "error registering user")
	}
	return u, nil
}
