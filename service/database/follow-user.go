package database

func (db *appdbimpl) FollowUser(uid string, fid string) error {
	_, queryError := db.c.Exec(FollowUser, fid, uid)
	if queryError != nil {
		db.logger.Debugf("Error following: %e", queryError)
		return queryError
	}
	return nil
}

func (db *appdbimpl) UnfollowUser(uid string, fid string) error {
	_, queryError := db.c.Exec(UnfollowUser, fid, uid)
	if queryError != nil {
		db.logger.Debugf("Error unfollowing: %e", queryError)
		return queryError
	}
	return nil
}
