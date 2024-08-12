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

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
	"github.com/sirupsen/logrus"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	LoginUser(string) (*schema.UserLogin, error)
	SearchUsersByName(string, string) ([]schema.SimpleUserData, error)
	SearchUidByToken(string) (string, error)
	ChangeUserName(string, string) error
	GetFollows(string) ([]schema.SimpleUserData, error)
	GetFollowers(string) ([]schema.SimpleUserData, error)
	FollowUser(string, string) error
	UnfollowUser(string, string) error
	GetBans(string) ([]schema.SimpleUserData, error)
	BanUser(string, string) error
	UnbanUser(string, string) error
	CreatePost(string) (string, error)
	GetSimplePost(string) (*schema.SimplePostData, error)
	GetMyStream(string) ([]schema.SimplePostData, error)
	GetUserProfile(string) (*schema.UserProfileData, error)
	GetUserPosts(string) ([]schema.SimplePostData, error)
	GetUserPost(string, string) (*schema.SimplePostData, error)
	DeletePhoto(string, string) error
	LikePost(string, string) error
	UnlikePost(string, string) error
	GetPostComments(string) ([]schema.PostComment, error)

	Ping() error
	Clean() error
}

type appdbimpl struct {
	c      *sql.DB
	logger *logrus.Logger
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
// Checks the existance of the important tables in the db.
// If one or more are missing and cannot be created it will return an error.
func New(db *sql.DB, logger *logrus.Logger) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	// Tables checks
	var tableName string
	for _, tableData := range TablesCheck {
		err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='` + tableData.TableName + `';`).Scan(&tableName)
		if errors.Is(err, sql.ErrNoRows) {
			_, err = db.Exec(tableData.CreateQuery)
			if err != nil {
				return nil, fmt.Errorf("error creating database table %s: %w", tableData.TableName, err)
			}
		}
	}

	return &appdbimpl{
		c:      db,
		logger: logger,
	}, nil
}

func (db *appdbimpl) Clean() error {
	for _, tableData := range TablesCheck {
		_, err := db.c.Exec("DELETE FROM " + tableData.TableName + ";")
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
