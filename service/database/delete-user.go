package database

import "github.com/SimoneDiCesare/WasaPhoto/service/database/queries"

func (db *appdbimpl) DeleteUser(uid string) (err error) {
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
	_, err = tx.Exec(queries.DeleteUser, uid)
	return err
}
