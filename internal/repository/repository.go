// Package repository contains Repository interface to interact with database.
package repository

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
)

// Repository interface is responsible for storing, retrieving and deleting data from database.
type Repository interface {
	UsersStorage
	BinaryFilesStorage
}

type BinaryFilesStorage interface {
	InsertBF(ctx context.Context, bf *binary_files.Files) error
	GetBF(ctx context.Context, userID, name string) (*binary_files.Files, error)
	UpdateBF(ctx context.Context, bf *binary_files.Files) (*binary_files.Files, error)
	DeleteBF(ctx context.Context, userID, name string) error
}

type LogPassesStorage interface {
	InsertLP(ctx context.Context, lp *log_passes.LogPasses) error
	GetLP(ctx context.Context, name string) (*log_passes.LogPasses, error)
	UpdateLP(ctx context.Context, lp *log_passes.LogPasses) (*log_passes.LogPasses, error)
	DeleteLP(ctx context.Context, name string) error
}

type CardsStorage interface {
	InsertC(ctx context.Context, c *cards.Cards) error
	GetC(ctx context.Context, name string) (*cards.Cards, error)
	UpdateC(ctx context.Context, c *cards.Cards) (*cards.Cards, error)
	DeleteC(ctx context.Context, name string) error
}

type TextsStorage interface {
	InsertT(ctx context.Context, t *texts.Texts) error
	GetT(ctx context.Context, name string) (*texts.Texts, error)
	UpdateT(ctx context.Context, t *texts.Texts) (*texts.Texts, error)
	DeleteT(ctx context.Context, name string) error
}

type UsersStorage interface {
	Register(ctx context.Context, credential *users.Credential) (*users.User, error)
	Login(ctx context.Context, login string) (cred *users.Credential, user *users.User, err error)
	InsertUser(ctx context.Context, u *users.User) error
}
