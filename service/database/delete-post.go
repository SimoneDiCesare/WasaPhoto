package database

import "github.com/SimoneDiCesare/WasaPhoto/service/database/queries"

func (db *appdbimpl) DeletePost(pid string, token string) (err error) {
	uid, uidError := db.GetUserIdFromToken(token)
	if uidError != nil {
		return uidError
	}
	// Start transaction
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			// Restore prevous state
			err = tx.Rollback()
		} else {
			// Commit changes
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec(queries.DeletePost, pid, uid)
	if err != nil {
		return err
	}
	return nil
}
