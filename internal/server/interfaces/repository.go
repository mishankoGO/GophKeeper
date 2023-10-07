// Package interfaces contains all the necessary interfaces to work with database.
// It has interfaces from binary files, cards, texts, users, credentials and logpasses.
package interfaces

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
	CardsStorage
	LogPassesStorage
	TextsStorage
}

// BinaryFilesStorage interface is responsible for storing, retrieving, updating and deleting binary files from database.
type BinaryFilesStorage interface {
	InsertBF(ctx context.Context, bf *binary_files.Files) error
	GetBF(ctx context.Context, userID, name string) (*binary_files.Files, error)
	UpdateBF(ctx context.Context, bf *binary_files.Files) (*binary_files.Files, error)
	DeleteBF(ctx context.Context, userID, name string) error
	ListBF(cts context.Context, userID string) ([]*binary_files.Files, error)
}

// LogPassesStorage interface is responsible for storing, retrieving, updating and deleting logpasses from database.

type LogPassesStorage interface {
	InsertLP(ctx context.Context, lp *log_passes.LogPasses) error
	GetLP(ctx context.Context, userID, name string) (*log_passes.LogPasses, error)
	UpdateLP(ctx context.Context, lp *log_passes.LogPasses) (*log_passes.LogPasses, error)
	DeleteLP(ctx context.Context, userID, name string) error
	ListLP(cts context.Context, userID string) ([]*log_passes.LogPasses, error)
}

// CardsStorage interface is responsible for storing, retrieving, updating and deleting cards from database.
type CardsStorage interface {
	InsertC(ctx context.Context, c *cards.Cards) error
	GetC(ctx context.Context, userID, name string) (*cards.Cards, error)
	UpdateC(ctx context.Context, c *cards.Cards) (*cards.Cards, error)
	DeleteC(ctx context.Context, userID, name string) error
	ListC(cts context.Context, userID string) ([]*cards.Cards, error)
}

// TextsStorage interface is responsible for storing, retrieving, updating and deleting texts from database.
type TextsStorage interface {
	InsertT(ctx context.Context, t *texts.Texts) error
	GetT(ctx context.Context, userID, name string) (*texts.Texts, error)
	UpdateT(ctx context.Context, t *texts.Texts) (*texts.Texts, error)
	DeleteT(ctx context.Context, userID, name string) error
	ListT(cts context.Context, userID string) ([]*texts.Texts, error)
}

// UsersStorage interface is responsible for registering, login and inserting user to database.
type UsersStorage interface {
	Register(ctx context.Context, credential *users.Credential) (*users.User, error)
	Login(ctx context.Context, login string) (cred *users.Credential, user *users.User, err error)
	InsertUser(ctx context.Context, u *users.User) error
}
