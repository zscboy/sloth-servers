package mysql

// errCode = 1 参数错误
func renameClub(clubID string, clubName string) (errCode int) {
	// dd_user_to_club`(IN userId varchar(64), IN clubId varchar(64))s

	stmt, err := dbConn.Prepare("call rename_club(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(clubID, clubName)
	err = row.Scan(&errCode)
	if err != nil {
		panic(err.Error())
	}

	return
}
