package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc/registration"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	query "github.com/mishankoGO/GophKeeper/internal/repository/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type DBRepository struct {
	DB *sql.DB
}

// NewDBRepository creates new repository instance
func NewDBRepository() (repository.Repository, error) {
	dataSourceName := "postgresql://gophkeeperuser:gophkeeperpwd@localhost:5432/gophkeeperdb?sslmode=disable"

	// open the connection to db
	DB, err := sql.Open("pgx", dataSourceName)
	if err != nil || DB == nil {
		return nil, fmt.Errorf("failed opening connection to %s: %w", dataSourceName, err)
	}

	// get the driver
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Printf("error getting the driver: %v", err)
		return nil, err
	}

	// go through migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Printf("error going through migrations: %v", err)
		return nil, err
	}

	err = m.Up()
	if err != nil {
		log.Printf("error doing up migrations: %v", err)
	}

	return &DBRepository{
		DB: DB,
	}, nil
}

func (r *DBRepository) Register(credential *pb.Credential) (*pb.User, error) {
	login, hashPassword := credential.Login, credential.HashPassword
	_, err := r.DB.Exec(query.RegisterQuery, login, hashPassword)
	if err != nil {
		return &pb.User{}, err
	}

	var userID string
	err = r.DB.QueryRow(query.GetUserId, login).Scan(&userID)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return &pb.User{}, status.Error(codes.InvalidArgument, fmt.Sprintf("no user with login %s", login))
	case err != nil:
		return &pb.User{}, status.Error(codes.Internal, "query error")
	}

	user := &pb.User{UserId: userID, CreatedAt: timestamppb.New(time.Now())}
	return user, nil
}


func