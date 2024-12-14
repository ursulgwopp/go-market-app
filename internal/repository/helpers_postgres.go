package


func (r *HelpersPostgres) CheckUsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}