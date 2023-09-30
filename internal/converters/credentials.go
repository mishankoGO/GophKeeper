// Package converters contains function to convert proto data to model data.
package converters

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

// CredentialToPBCredential converts model credential to proto credential.
func CredentialToPBCredential(c *users.Credential) *pb.Credential {
	return &pb.Credential{Login: c.Login, Password: c.Password}
}

// PBCredentialToCredential converts proto credential to model credential.
func PBCredentialToCredential(pbc *pb.Credential) *users.Credential {
	return &users.Credential{Login: pbc.Login, Password: pbc.Password}
}
