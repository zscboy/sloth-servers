package mysql

// errCode = 1 参数错误， 2 已经在牌友群，3牌友群不存在
func changeClubMemberRole(memberID string, clubID string, role int32) (errCode int) {
	// `add_user_to_club`(IN userId varchar(64), IN clubId varchar(64), IN clubRole int)

	stmt, err := dbConn.Prepare("call change_club_member_role(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(memberID, clubID, role)
	err = row.Scan(&errCode)
	if err != nil {
		panic(err.Error())
	}

	return
}
