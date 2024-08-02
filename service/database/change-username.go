package database

func (db *appdbimpl) ChangeUserName(uid string, username string) error {
	result, err := db.c.Exec(UpdateUserName, username, uid)
	if err != nil {
		db.logger.Debugf("Error on db Query: %w", err)
		return err
	}
	count, err := result.RowsAffected()
	if count == 0 || err != nil {
		db.logger.Debugf("No rows affected: %d, %w", count, err)
		return err
	}
	return nil
}
