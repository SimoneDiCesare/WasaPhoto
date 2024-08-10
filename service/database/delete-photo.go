package database

func (db *appdbimpl) DeletePhoto(uid string, pid string) error {
	_, queryError := db.c.Exec(DeletePhoto, pid, uid)
	if queryError != nil {
		db.logger.Debugf("Error deleting photo: %e", queryError)
		return queryError
	}
	return nil
}
