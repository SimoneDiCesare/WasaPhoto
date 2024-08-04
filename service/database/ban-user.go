package database

func (db *appdbimpl) BanUser(uid string, bid string) error {
	_, queryError := db.c.Exec(BanUser, uid, bid)
	if queryError != nil {
		db.logger.Debugf("Error banning: %w", queryError)
		return queryError
	}
	return nil
}

func (db *appdbimpl) UnbanUser(uid string, bid string) error {
	_, queryError := db.c.Exec(UnbanUser, uid, bid)
	if queryError != nil {
		db.logger.Debugf("Error ubanning: %w", queryError)
		return queryError
	}
	return nil
}
