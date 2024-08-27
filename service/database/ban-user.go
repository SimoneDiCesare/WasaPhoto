package database

func (db *appdbimpl) BanUser(uid string, bid string) error {
	// First remove follows/follower tuples from db
	err := db.UnfollowUser(uid, bid)
	if err != nil {
		db.logger.Errorf("Can't unfollow while banning: %e", err)
		return err
	}
	err = db.UnfollowUser(bid, uid)
	if err != nil {
		db.logger.Errorf("Can't unfollow while banning: %e", err)
		return err
	}
	_, queryError := db.c.Exec(BanUser, uid, bid)
	if queryError != nil {
		db.logger.Debugf("Error banning: %e", queryError)
		return queryError
	}
	return nil
}

func (db *appdbimpl) UnbanUser(uid string, bid string) error {
	_, queryError := db.c.Exec(UnbanUser, uid, bid)
	if queryError != nil {
		db.logger.Debugf("Error ubanning: %e", queryError)
		return queryError
	}
	return nil
}
