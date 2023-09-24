package converters

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

func CredentialToPBCredential(c *users.Credential) *pb.Credential {
	return &pb.Credential{Login: c.Login, HashPassword: c.HashPassword}
}

func PBCredentialToCredential(pbc *pb.Credential) *users.Credential {
	return &users.Credential{Login: pbc.Login, HashPassword: pbc.HashPassword}
}
