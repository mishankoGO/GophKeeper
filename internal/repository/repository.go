// Package repository contains Repository interface to interact with database.
package repository

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc/registration"
)

// Repository interface is responsible for storing, retrieving and deleting data from database.
type Repository interface {
	Register(credential *pb.Credential) (*pb.User, error)
}

type LogPassesStorage interface {
}

type CardsStorage interface {
}

type TextsStorage interface {
}

type UsersStorage interface {
	Insert()
}
