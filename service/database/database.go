/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
	"github.com/sirupsen/logrus"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	LoginUser(username string) (int, User, error)
	VerifyToken(token string) error
	GetUserProfile(uid string) (*UserProfile, error)
	SearchUsers(token string, text string) ([]SimpleUserProfile, error)
	VerifyUidToken(uid string, token string) error
	ChangeUserName(username string, uid string) error
	DeleteUser(uid string) error
	GetFollows(uid string) ([]SimpleUserProfile, error)
	FollowUser(uid1 string, uid2 string) error
	UnfollowUser(uid1 string, uid2 string) error
	GetFollowers(uid string) ([]SimpleUserProfile, error)
	BanUser(uid1 string, uid2 string) error
	UnbanUser(uid1 string, uid2 string) error
	GetBans(uid string) ([]SimpleUserProfile, error)
	CreatePost(uid string, caption string) (*Post, error)
	GetPost(pid string) (*Post, error)

	Ping() error
}

type appdbimpl struct {
	c      *sql.DB
	logger *logrus.Logger
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB, logger *logrus.Logger) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	// Check tables in db, and creates it if they are missing
	logger.Debugf("Checking db integrity ")
	for _, tableCheck := range queries.TablesCheck {
		var flag string
		logger.Debugf("Checking table %s", tableCheck.TableName)
		err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tableCheck.TableName).Scan(&flag)
		if errors.Is(err, sql.ErrNoRows) {
			err = createTable(db, tableCheck.CreateQuery, logger)
			if err != nil {
				return nil, fmt.Errorf("error creating database structure: %w", err)
			}
		} else {
			logger.Debugf("Table %s OK", tableCheck.TableName)
		}
	}

	return &appdbimpl{
		c:      db,
		logger: logger,
	}, nil
}

func createTable(db *sql.DB, query string, logger *logrus.Logger) error {
	logger.Debugf("Creating table:\n\t%s\n", query)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GetImage(path string) string {
	if fileExists(path) {
		return path
	}
	if strings.Contains(path, "users") {
		return "/defaults/profile_default.png"
	} else {
		return "/defaults/empty.png"
	}
}
