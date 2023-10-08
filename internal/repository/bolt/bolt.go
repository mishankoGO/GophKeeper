package bolt

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"strings"
	"time"
)

type DBRepository struct {
	conf *config.Config
	DB   *bolt.DB
}

func NewDBRepository(conf *config.Config) (interfaces.Repository, error) {
	name := fmt.Sprintf("internal/repository/bolt/backup/gophKeeper.db")
	db, err := bolt.Open(name, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error connecting to bolt db: %w", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("Credentials"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %w", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("BinaryFiles"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %w", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("Texts"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %w", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("Cards"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %w", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("LogPasses"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error updating db: %w", err)
	}
	return &DBRepository{conf: conf, DB: db}, nil
}

func (r *DBRepository) Close() error {
	return r.DB.Close()
}

func (r *DBRepository) InsertUser(cred *users.Credential, user *users.User) error {
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Credentials"))

		err := b.Put([]byte("user_id"), []byte(user.UserID))
		if err != nil {
			return err
		}
		err = b.Put([]byte("login"), []byte(cred.Login))
		if err != nil {
			return err
		}
		err = b.Put([]byte("password"), []byte(cred.Password))
		if err != nil {
			return err
		}
		err = b.Put([]byte("created_at"), []byte(user.CreatedAt.String()))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error inserting new user: %w", err)
	}
	return nil
}

// Login method is responsible for retrieving userID from database.
func (r *DBRepository) Login(login string) (*users.Credential, *users.User, error) {

	// login user
	var log, password, userId string
	var createdAtt time.Time
	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Credentials"))

		uid := b.Get([]byte("user_id"))
		l := b.Get([]byte("login"))
		p := b.Get([]byte("password"))
		c := b.Get([]byte("created_at"))
		if l == nil || p == nil || string(l) != login {
			return fmt.Errorf("user is not registered")
		}
		log, password, userId = string(l), string(p), string(uid)

		format := "2006-01-02 15:04:05"
		createdAt, err := time.Parse(format, strings.Split(string(c), ".")[0])
		if err != nil {
			return fmt.Errorf("error parsing creation time: %w", err)
		}

		createdAtt = createdAt

		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error retrieving login info: %w", err)
	}

	// create user and credential
	var cred = &users.Credential{Login: log, Password: password}
	var user = &users.User{UserID: userId, Login: login, CreatedAt: createdAtt}

	return cred, user, nil
}

// InsertBF method inserts binary file to db.
func (r *DBRepository) InsertBF(bf *binary_files.Files) error {
	// insert binary file
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BinaryFiles"))

		bfBytes, err := json.Marshal(bf)
		if err != nil {
			return fmt.Errorf("error marshalling binary file: %w", err)
		}

		err = b.Put([]byte(bf.Name), bfBytes)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error inserting new binary file: %w", err)
	}
	return nil
}

// GetBF method retrieves binary file from db.
func (r *DBRepository) GetBF(name string) (*binary_files.Files, error) {
	var bf binary_files.Files

	// get binary file by name
	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BinaryFiles"))
		v := b.Get([]byte(name))
		err := json.Unmarshal(v, &bf)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error getting binary file %s: %w", name, err)
	}

	return &bf, nil
}

// UpdateBF method updates binary file.
func (r *DBRepository) UpdateBF(bf *binary_files.Files) (*binary_files.Files, error) {
	// update binary file with meta
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BinaryFiles"))

		bfBytes, err := json.Marshal(bf)
		if err != nil {
			return fmt.Errorf("error marshalling binary file: %w", err)
		}

		err = b.Put([]byte(bf.Name), bfBytes)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error updating binary file: %w", err)
	}

	return bf, nil
}

// DeleteBF method deletes binary file from db.
func (r *DBRepository) DeleteBF(name string) error {
	// delete binary file
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BinaryFiles"))
		err := b.Delete([]byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error deleting binary file %s: %w", name, err)
	}
	return nil
}

// ListBF method returns list of binary files.
func (r *DBRepository) ListBF() ([]*binary_files.Files, error) {
	var bfs []*binary_files.Files
	err := r.DB.View(func(tx *bolt.Tx) error {
		var bf binary_files.Files
		b := tx.Bucket([]byte("BinaryFiles"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &bf)
			if err != nil {
				return err
			}
			bfs = append(bfs, &bf)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error listing binary files: %w", err)
	}
	return bfs, nil
}

// InsertLP method inserts log pass to db.
func (r *DBRepository) InsertLP(lp *log_passes.LogPasses) error {
	// insert binary file
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("LogPasses"))

		lpBytes, err := json.Marshal(lp)
		if err != nil {
			return fmt.Errorf("error marshalling log pass: %w", err)
		}

		err = b.Put([]byte(lp.Name), lpBytes)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error inserting new log pass: %w", err)
	}
	return nil
}

// GetLP method retrieves log pass from db.
func (r *DBRepository) GetLP(name string) (*log_passes.LogPasses, error) {
	var lp log_passes.LogPasses

	// get binary file by name
	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("LogPasses"))
		v := b.Get([]byte(name))
		err := json.Unmarshal(v, &lp)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error getting log pass %s: %w", name, err)
	}

	return &lp, nil
}

// UpdateLP method updates log pass in db.
func (r *DBRepository) UpdateLP(lp *log_passes.LogPasses) (*log_passes.LogPasses, error) {
	// update log pass with meta
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("LogPasses"))

		lpBytes, err := json.Marshal(lp)
		if err != nil {
			return fmt.Errorf("error marshalling log pass: %w", err)
		}

		err = b.Put([]byte(lp.Name), lpBytes)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error updating log pass: %w", err)
	}

	return lp, nil
}

// DeleteLP method deletes log pass from db.
func (r *DBRepository) DeleteLP(name string) error {
	// delete log pass
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("LogPasses"))
		err := b.Delete([]byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error deleting log pass %s: %w", name, err)
	}
	return nil
}

// ListLP method returns list of log passes.
func (r *DBRepository) ListLP() ([]*log_passes.LogPasses, error) {
	var lps []*log_passes.LogPasses
	err := r.DB.View(func(tx *bolt.Tx) error {
		var lp log_passes.LogPasses
		b := tx.Bucket([]byte("LogPasses"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &lp)
			if err != nil {
				return err
			}
			lps = append(lps, &lp)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error listing logpasses: %w", err)
	}
	return lps, nil
}

// InsertC method inserts card to db.
func (r *DBRepository) InsertC(c *cards.Cards) error {
	// insert card
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Cards"))

		cBytes, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("error marshalling card: %w", err)
		}

		err = b.Put([]byte(c.Name), cBytes)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error inserting new card: %w", err)
	}
	return nil
}

// GetC method retrieves card from db.
func (r *DBRepository) GetC(name string) (*cards.Cards, error) {
	var c cards.Cards

	// get binary file by name
	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Cards"))
		v := b.Get([]byte(name))
		err := json.Unmarshal(v, &c)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error getting card %s: %w", name, err)
	}

	return &c, nil
}

// UpdateC method updates card in db.
func (r *DBRepository) UpdateC(c *cards.Cards) (*cards.Cards, error) {
	// update card with meta
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Cards"))

		cBytes, err := json.Marshal(c)
		if err != nil {
			return fmt.Errorf("error marshalling card: %w", err)
		}

		err = b.Put([]byte(c.Name), cBytes)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error updating card: %w", err)
	}

	return c, nil
}

// DeleteC method deletes card from db.
func (r *DBRepository) DeleteC(name string) error {
	// delete card
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Cards"))
		err := b.Delete([]byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error deleting card %s: %w", name, err)
	}
	return nil
}

// ListC method returns all the cards.
func (r *DBRepository) ListC() ([]*cards.Cards, error) {
	var ccs []*cards.Cards
	err := r.DB.View(func(tx *bolt.Tx) error {
		var cc cards.Cards
		b := tx.Bucket([]byte("Cards"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &cc)
			if err != nil {
				return err
			}
			ccs = append(ccs, &cc)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error listing cards: %w", err)
	}
	return ccs, nil
}

// InsertT inserts text in db.
func (r *DBRepository) InsertT(t *texts.Texts) error {
	// insert text
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Texts"))

		tBytes, err := json.Marshal(t)
		if err != nil {
			return fmt.Errorf("error marshalling text: %w", err)
		}

		err = b.Put([]byte(t.Name), tBytes)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error inserting new text: %w", err)
	}
	return nil
}

// GetT retrieves text from db.
func (r *DBRepository) GetT(name string) (*texts.Texts, error) {
	var t texts.Texts

	// get binary file by name
	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Texts"))
		v := b.Get([]byte(name))
		err := json.Unmarshal(v, &t)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error getting text %s: %w", name, err)
	}

	return &t, nil
}

// UpdateT method updates text in db.
func (r *DBRepository) UpdateT(t *texts.Texts) (*texts.Texts, error) {
	// update text with meta
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Texts"))

		tBytes, err := json.Marshal(t)
		if err != nil {
			return fmt.Errorf("error marshalling text: %w", err)
		}

		err = b.Put([]byte(t.Name), tBytes)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error updating text: %w", err)
	}

	return t, nil
}

// DeleteT deletes text from db.
func (r *DBRepository) DeleteT(name string) error {
	// delete text
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Texts"))
		err := b.Delete([]byte(name))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error deleting text %s: %w", name, err)
	}
	return nil
}

// ListT method returns all the texts.
func (r *DBRepository) ListT() ([]*texts.Texts, error) {
	var ts []*texts.Texts
	err := r.DB.View(func(tx *bolt.Tx) error {
		var t texts.Texts
		b := tx.Bucket([]byte("Texts"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
			ts = append(ts, &t)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error listing texts: %w", err)
	}
	return ts, nil
}
