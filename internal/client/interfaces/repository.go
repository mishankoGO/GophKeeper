// Package interfaces contains all the necessary interfaces to work with database.
// It has interfaces from binary files, cards, texts, users, credentials and logpasses.
package interfaces

import (
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
	InsertBF(bf *binary_files.Files) error
	GetBF(name string) (*binary_files.Files, error)
	UpdateBF(bf *binary_files.Files) (*binary_files.Files, error)
	DeleteBF(name string) error
	ListBF() ([]*binary_files.Files, error)
}

// LogPassesStorage interface is responsible for storing, retrieving, updating and deleting logpasses from database.

type LogPassesStorage interface {
	InsertLP(lp *log_passes.LogPasses) error
	GetLP(name string) (*log_passes.LogPasses, error)
	UpdateLP(lp *log_passes.LogPasses) (*log_passes.LogPasses, error)
	DeleteLP(name string) error
	ListLP() ([]*log_passes.LogPasses, error)
}

// CardsStorage interface is responsible for storing, retrieving, updating and deleting cards from database.
type CardsStorage interface {
	InsertC(c *cards.Cards) error
	GetC(name string) (*cards.Cards, error)
	UpdateC(c *cards.Cards) (*cards.Cards, error)
	DeleteC(name string) error
	ListC() ([]*cards.Cards, error)
}

// TextsStorage interface is responsible for storing, retrieving, updating and deleting texts from database.
type TextsStorage interface {
	InsertT(t *texts.Texts) error
	GetT(name string) (*texts.Texts, error)
	UpdateT(t *texts.Texts) (*texts.Texts, error)
	DeleteT(name string) error
	ListT() ([]*texts.Texts, error)
}

// UsersStorage interface is responsible for registering, login and inserting user to database.
type UsersStorage interface {
	Login(login string) (cred *users.Credential, user *users.User, err error)
	InsertUser(cred *users.Credential, user *users.User) error
	Close() error
}
