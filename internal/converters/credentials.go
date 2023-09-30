package converters

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

func CredentialToPBCredential(c *users.Credential) *pb.Credential {
	return &pb.Credential{Login: c.Login, Password: c.Password}
}

func PBCredentialToCredential(pbc *pb.Credential) *users.Credential {
	return &users.Credential{Login: pbc.Login, Password: pbc.Password}
}
