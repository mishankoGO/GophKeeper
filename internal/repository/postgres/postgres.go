// Package postgres offers a functionality to work with Postgres database.
// It can insert, read, delete and update data.
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	query "github.com/mishankoGO/GophKeeper/internal/repository/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

// DBRepository contains database handle.
type DBRepository struct {
	DB *sql.DB
}

// NewDBRepository creates new repository instance.
func NewDBRepository(conf *config.Config) (repository.Repository, error) {
	//dataSourceName := "postgresql://gophkeeperuser:gophkeeperpwd@localhost:5432/gophkeeperdb?sslmode=disable"

	// get db dsn from config
	dataSourceName := conf.DatabaseDSN

	// open the connection to db
	DB, err := sql.Open("pgx", dataSourceName)
	if err != nil || DB == nil {
		return nil, fmt.Errorf("failed opening connection to %s: %w", dataSourceName, err)
	}

	// get the driver
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("error getting the database driver: %w", err)
	}

	// go through migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("error creating migrations: %w", err)
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("migrations are all up")
	}

	return &DBRepository{
		DB: DB,
	}, nil
}

// Register method is responsible for creating new credential and user record in the database.
func (r *DBRepository) Register(ctx context.Context, credential *users.Credential) (user *users.User, err error) {

	// get login and password from input
	login, hashPassword := credential.Login, credential.HashPassword

	// check if user exists
	_, err = r.DB.ExecContext(ctx, query.CheckUser, login)
	if err != nil {
		return user, status.Errorf(codes.AlreadyExists, "user %s already exists: %v", login, err)
	}

	// insert new credential in db
	_, err = r.DB.ExecContext(ctx, query.RegisterQuery, login, hashPassword)
	if err != nil {
		return user, status.Errorf(codes.Internal, "error inserting new credential: %v", err)
	}

	// get user id
	var unitID string
	err = r.DB.QueryRowContext(ctx, query.GetUserId, login).Scan(&unitID)
	if err != nil {
		return user, status.Errorf(codes.Internal, "error getting %s id: %v", login, err)
	}

	// create user instance
	user = &users.User{UserID: unitID, Login: login, CreatedAt: time.Now()}

	// insert new user
	err = r.InsertUser(ctx, user)
	if err != nil {
		return user, status.Errorf(codes.Internal, "error inserting new user %v: %v", user, err)
	}

	return user, nil
}

// Login method is responsible for retrieving userID from database.
func (r *DBRepository) Login(ctx context.Context, login string) (*users.Credential, *users.User, error) {

	var userID, hashPassword string
	var createdAt time.Time
	err := r.DB.QueryRowContext(ctx, query.LoginUser, login).Scan(&userID, &hashPassword, &createdAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil, status.Errorf(codes.InvalidArgument, "no user with login %s", login)
	case err != nil:
		return nil, nil, status.Error(codes.Internal, "error getting user id")
	}

	var cred = &users.Credential{Login: login, HashPassword: hashPassword}
	var user = &users.User{UserID: userID, Login: login, CreatedAt: createdAt}

	return cred, user, nil
}

// InsertUser method is responsible for inserting new user to users table.
func (r *DBRepository) InsertUser(ctx context.Context, u *users.User) error {

	// insert new user
	_, err := r.DB.ExecContext(ctx, query.AddUserQuery, u.UserID, u.Login, u.CreatedAt)
	if err != nil {
		return status.Error(codes.Internal, "error inserting new user")
	}
	return nil
}

func (r *DBRepository) InsertBF(ctx context.Context, bf *binary_files.Files) error {
	return nil
}

func (r *DBRepository) GetBF(ctx context.Context, name string) (*binary_files.Files, error) {
	return nil, nil
}

func (r *DBRepository) UpdateBF(ctx context.Context, bf *binary_files.Files) (*binary_files.Files, error) {
	return nil, nil
}

func (r *DBRepository) DeleteBF(ctx context.Context, name string) error {
	return nil
}

func (r *DBRepository) InsertLP(ctx context.Context, lp *log_passes.LogPasses) error {
	return nil
}

func (r *DBRepository) GetLP(ctx context.Context, name string) (*log_passes.LogPasses, error) {
	return nil, nil
}

func (r *DBRepository) UpdateLP(ctx context.Context, lp *log_passes.LogPasses) (*log_passes.LogPasses, error) {
	return nil, nil
}

func (r *DBRepository) DeleteLP(ctx context.Context, name string) error {
	return nil
}

func (r *DBRepository) InsertC(ctx context.Context, c *cards.Cards) error {
	return nil
}

func (r *DBRepository) GetC(ctx context.Context, name string) (*cards.Cards, error) {
	return nil, nil
}

func (r *DBRepository) UpdateC(ctx context.Context, c *cards.Cards) (*cards.Cards, error) {
	return nil, nil
}

func (r *DBRepository) DeleteC(ctx context.Context, name string) error {
	return nil
}

func (r *DBRepository) InsertT(ctx context.Context, t *texts.Texts) error {
	return nil
}

func (r *DBRepository) GetT(ctx context.Context, name string) (*texts.Texts, error) {
	return nil, nil
}

func (r *DBRepository) UpdateT(ctx context.Context, t *texts.Texts) (*texts.Texts, error) {
	return nil, nil
}

func (r *DBRepository) DeleteT(ctx context.Context, name string) error {
	return nil
}
