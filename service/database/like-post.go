package database

import "github.com/SimoneDiCesare/WasaPhoto/service/database/queries"

func (db *appdbimpl) LikePost(pid string, token string) error {
	uid, getError := db.GetUserIdFromToken(token)
	if getError != nil {
		return getError
	}
	r, err := db.c.Exec(queries.AddLikeToPost, pid, uid)
	if err != nil {
		return err
	}
	_, err = r.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
