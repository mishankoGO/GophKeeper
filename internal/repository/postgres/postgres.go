// Package postgres offers a functionality to work with Postgres database.
// It can insert, read, delete and update data.
package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
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
	query "github.com/mishankoGO/GophKeeper/internal/repository/sql"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"log"
	"time"
)

// DBRepository contains database handle.
type DBRepository struct {
	DB *sql.DB
}

// NewDBRepository creates new repository instance.
func NewDBRepository(conf *config.Config) (interfaces.Repository, error) {
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
	login, Password := credential.Login, credential.Password

	// check if user exists
	_, err = r.DB.ExecContext(ctx, query.CheckUser, login)
	if err != nil {
		return user, fmt.Errorf("user %s already exists: %w", login, err)
	}

	// insert new credential in db
	_, err = r.DB.ExecContext(ctx, query.RegisterQuery, login, Password)
	if err != nil {
		return user, fmt.Errorf("error inserting new credential: %w", err)
	}

	// get user id
	var unitID string
	err = r.DB.QueryRowContext(ctx, query.GetUserId, login).Scan(&unitID)
	if err != nil {
		return user, fmt.Errorf("error getting %s id: %w", login, err)
	}

	// create user instance
	user = &users.User{UserID: unitID, Login: login, CreatedAt: time.Now()}

	// insert new user
	err = r.InsertUser(ctx, user)
	if err != nil {
		return user, fmt.Errorf("error inserting new user %v: %w", user, err)
	}

	return user, nil
}

// Login method is responsible for retrieving userID from database.
func (r *DBRepository) Login(ctx context.Context, login string) (*users.Credential, *users.User, error) {

	var userID, Password string
	var createdAt time.Time

	// login user
	err := r.DB.QueryRowContext(ctx, query.LoginUser, login).Scan(&userID, &Password, &createdAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil, fmt.Errorf("no user with login %s", login)
	case err != nil:
		return nil, nil, fmt.Errorf("error getting user id: %w", err)
	}

	// create user and credential
	var cred = &users.Credential{Login: login, Password: Password}
	var user = &users.User{UserID: userID, Login: login, CreatedAt: createdAt}

	return cred, user, nil
}

// InsertUser method is responsible for inserting new user to users table.
func (r *DBRepository) InsertUser(ctx context.Context, u *users.User) error {
	// insert new user
	_, err := r.DB.ExecContext(ctx, query.AddUserQuery, u.UserID, u.Login, u.CreatedAt)
	if err != nil {
		return fmt.Errorf("error inserting new user: %w", err)
	}
	return nil
}

// InsertBF method inserts binary file to db.
func (r *DBRepository) InsertBF(ctx context.Context, bf *binary_files.Files) error {
	// insert binary file
	_, err := r.DB.ExecContext(ctx, query.InsertBinaryFile, bf.UserID, bf.Name, bf.File, bf.UpdatedAt, bf.Meta)
	if err != nil {
		return fmt.Errorf("error inserting new binary file: %w", err)
	}
	return nil
}

// GetBF method retrieves binary file from db.
func (r *DBRepository) GetBF(ctx context.Context, userID, name string) (*binary_files.Files, error) {
	var uid, n string
	var updatedAt time.Time
	var meta, file []byte

	// get binary file by name
	err := r.DB.QueryRowContext(ctx, query.GetBinaryFile, userID, name).Scan(&uid, &n, &file, &updatedAt, &meta)
	if err != nil {
		return nil, fmt.Errorf("error getting binary file %s: %w", n, err)
	}

	// unmarshall metadata
	var metaMap = make(map[string]string)
	err = json.Unmarshal(meta, &metaMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling meta information: %w", err)
	}

	// create binary file
	var bf = &binary_files.Files{UserID: uid, Name: n, File: file, UpdatedAt: updatedAt, Meta: metaMap}

	return bf, nil
}

// UpdateBF method updates binary file.
func (r *DBRepository) UpdateBF(ctx context.Context, bf *binary_files.Files) (*binary_files.Files, error) {
	// marshall metadata if present
	if bf.Meta != nil {
		metaByte, err := json.Marshal(bf.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling meta map: %w", err)
		}

		// update binary file with meta
		_, err = r.DB.ExecContext(ctx, query.UpdateBinaryFile, bf.UserID, bf.Name, bf.File, bf.UpdatedAt, metaByte)
		if err != nil {
			return nil, fmt.Errorf("error updating binary file: %w", err)
		}
	}

	// update binary file without metadata
	_, err := r.DB.ExecContext(ctx, query.UpdateBinaryFile, bf.UserID, bf.Name, bf.File, bf.UpdatedAt, bf.Meta)
	if err != nil {
		return nil, fmt.Errorf("error updating binary file: %w", err)
	}

	return bf, nil
}

// DeleteBF method deletes binary file from db.
func (r *DBRepository) DeleteBF(ctx context.Context, userID, name string) error {
	// delete binary file
	_, err := r.DB.ExecContext(ctx, query.DeleteBinaryFile, userID, name)
	if err != nil {
		return fmt.Errorf("error deleting binary file %s: %w", name, err)
	}
	return nil
}

// InsertLP method inserts log pass to db.
func (r *DBRepository) InsertLP(ctx context.Context, lp *log_passes.LogPasses) error {
	// insert log pass
	_, err := r.DB.ExecContext(ctx, query.InsertLogPass, lp.UserID, lp.Name, lp.Login, lp.Password, lp.UpdatedAt, lp.Meta)
	if err != nil {
		return fmt.Errorf("error inserting new log pass: %w", err)
	}
	return nil
}

// GetLP method retrieves log pass from db.
func (r *DBRepository) GetLP(ctx context.Context, userID, name string) (*log_passes.LogPasses, error) {
	var uid, n string
	var updatedAt time.Time
	var meta, login, password []byte

	// get log pass by name
	err := r.DB.QueryRowContext(ctx, query.GetLogPass, userID, name).Scan(&uid, &n, &login, &password, &updatedAt, &meta)
	if err != nil {
		return nil, fmt.Errorf("error getting log pass %s: %w", n, err)
	}

	// unmarshall metadata
	var metaMap = make(map[string]string)
	err = json.Unmarshal(meta, &metaMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling meta information: %w", err)
	}

	// create log pass
	var lp = &log_passes.LogPasses{UserID: uid, Name: n, Login: login, Password: password, UpdatedAt: updatedAt, Meta: metaMap}

	return lp, nil
}

// UpdateLP method updates log pass in db.
func (r *DBRepository) UpdateLP(ctx context.Context, lp *log_passes.LogPasses) (*log_passes.LogPasses, error) {
	// marshall meta if present
	if lp.Meta != nil {
		metaByte, err := json.Marshal(lp.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling meta map: %w", err)
		}

		// update log pass
		_, err = r.DB.ExecContext(ctx, query.UpdateLogPass, lp.UserID, lp.Name, lp.Login, lp.Password, lp.UpdatedAt, metaByte)
		if err != nil {
			return nil, fmt.Errorf("error updating log pass: %w", err)
		}
	}

	// update log pass without meta
	_, err := r.DB.ExecContext(ctx, query.UpdateLogPass, lp.UserID, lp.Name, lp.Login, lp.Password, lp.UpdatedAt, lp.Meta)
	if err != nil {
		return nil, fmt.Errorf("error updating log pass: %w", err)
	}

	return lp, nil
}

// DeleteLP method deletes log pass from db.
func (r *DBRepository) DeleteLP(ctx context.Context, userID, name string) error {
	_, err := r.DB.ExecContext(ctx, query.DeleteLogPass, userID, name)
	if err != nil {
		return fmt.Errorf("error deleting log pass %s: %w", name, err)
	}
	return nil
}

// InsertC method inserts card to db.
func (r *DBRepository) InsertC(ctx context.Context, c *cards.Cards) error {
	// insert card
	_, err := r.DB.ExecContext(ctx, query.InsertCard, c.UserID, c.Name, c.Card, c.UpdatedAt, c.Meta)
	if err != nil {
		return fmt.Errorf("error inserting new card: %w", err)
	}
	return nil
}

// GetC method retrieves card from db.
func (r *DBRepository) GetC(ctx context.Context, userID, name string) (*cards.Cards, error) {
	// get card by name
	var uid, n string
	var updatedAt time.Time
	var meta, card []byte

	err := r.DB.QueryRowContext(ctx, query.GetCard, userID, name).Scan(&uid, &n, &card, &updatedAt, &meta)
	if err != nil {
		return nil, fmt.Errorf("error getting card %s: %v", n, err)
	}

	var metaMap = make(map[string]string)
	err = json.Unmarshal(meta, &metaMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling meta information: %v", err)
	}

	var c = &cards.Cards{UserID: uid, Name: n, Card: card, UpdatedAt: updatedAt, Meta: metaMap}
	return c, nil
}

// UpdateC method updates card in db.
func (r *DBRepository) UpdateC(ctx context.Context, c *cards.Cards) (*cards.Cards, error) {
	if c.Meta != nil {
		metaByte, err := json.Marshal(c.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling meta map: %v", err)
		}
		_, err = r.DB.ExecContext(ctx, query.UpdateCard, c.UserID, c.Name, c.Card, c.UpdatedAt, metaByte)
		if err != nil {
			return nil, fmt.Errorf("error updating card: %v", err)
		}
	}
	_, err := r.DB.ExecContext(ctx, query.UpdateCard, c.UserID, c.Name, c.Card, c.UpdatedAt, c.Meta)
	if err != nil {
		return nil, fmt.Errorf("error updating card: %v", err)
	}

	return c, nil
}

// DeleteC method deletes card from db.
func (r *DBRepository) DeleteC(ctx context.Context, userID, name string) error {
	_, err := r.DB.ExecContext(ctx, query.DeleteCard, userID, name)
	if err != nil {
		return fmt.Errorf("error deleting card %s: %v", name, err)
	}
	return nil
}

// InsertT inserts text in db.
func (r *DBRepository) InsertT(ctx context.Context, t *texts.Texts) error {
	// insert text
	_, err := r.DB.ExecContext(ctx, query.InsertText, t.UserID, t.Name, t.Text, t.UpdatedAt, t.Meta)
	if err != nil {
		return fmt.Errorf("error inserting new text: %v", err)
	}
	return nil
}

// GetT retrieves text from db.
func (r *DBRepository) GetT(ctx context.Context, userID, name string) (*texts.Texts, error) {
	// get card by name
	var uid, n string
	var updatedAt time.Time
	var meta, text []byte

	err := r.DB.QueryRowContext(ctx, query.GetText, userID, name).Scan(&uid, &n, &text, &updatedAt, &meta)
	if err != nil {
		return nil, fmt.Errorf("error getting text %s: %v", n, err)
	}

	var metaMap = make(map[string]string)
	err = json.Unmarshal(meta, &metaMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling meta information: %v", err)
	}

	var t = &texts.Texts{UserID: uid, Name: n, Text: text, UpdatedAt: updatedAt, Meta: metaMap}
	return t, nil
}

// UpdateT method updates text in db.
func (r *DBRepository) UpdateT(ctx context.Context, t *texts.Texts) (*texts.Texts, error) {
	if t.Meta != nil {
		metaByte, err := json.Marshal(t.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling meta map: %v", err)
		}
		_, err = r.DB.ExecContext(ctx, query.UpdateText, t.UserID, t.Name, t.Text, t.UpdatedAt, metaByte)
		if err != nil {
			return nil, fmt.Errorf("error updating text: %v", err)
		}
	}
	_, err := r.DB.ExecContext(ctx, query.UpdateText, t.UserID, t.Name, t.Text, t.UpdatedAt, t.Meta)
	if err != nil {
		return nil, fmt.Errorf("error updating text: %v", err)
	}

	return t, nil
}

// DeleteT deletes text from db.
func (r *DBRepository) DeleteT(ctx context.Context, userID, name string) error {
	_, err := r.DB.ExecContext(ctx, query.DeleteText, userID, name)
	if err != nil {
		return fmt.Errorf("error deleting text %s: %v", name, err)
	}
	return nil
}
