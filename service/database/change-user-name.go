package database

import (
	"errors"
	"fmt"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) ChangeUserName(username string, uid string) error {
	result, err := db.c.Exec(queries.ChangeUsername, username, uid)
	if err != nil {
		return fmt.Errorf("failed to update username: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no user found")
	}
	return nil
}
