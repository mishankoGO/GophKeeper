package repository

import pb "github.com/mishankoGO/GophKeeper/api"

type Repository interface {
	Register(credential *pb.Credential) (*pb.User, error)
}
