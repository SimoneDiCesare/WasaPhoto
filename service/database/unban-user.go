package database

import "github.com/SimoneDiCesare/WasaPhoto/service/database/queries"

func (db *appdbimpl) UnbanUser(uid1 string, uid2 string) error {
	r, err := db.c.Exec(queries.UnbanUser, uid1, uid2)
	if err != nil {
		return err
	}
	_, err = r.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
