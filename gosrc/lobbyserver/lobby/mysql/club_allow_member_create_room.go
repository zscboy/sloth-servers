package mysql

// errCode = 1 参数错误， 2 已经在牌友群，3牌友群不存在
func allowMemberCreateRoom(memberID string, clubID string, permission int32) (errCode int) {
	// `add_user_to_club`(IN userId varchar(64), IN clubId varchar(64), IN clubRole int)

	stmt, err := dbConn.Prepare("call change_club_member_permission(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(memberID, clubID, permission)
	err = row.Scan(&errCode)
	if err != nil {
		panic(err.Error())
	}

	return
}
