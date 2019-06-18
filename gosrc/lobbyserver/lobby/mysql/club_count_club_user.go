package mysql

// 统计用户的牌友群数量
func countClubUser(clubID string) (count int) {
	stmt, err := dbConn.Prepare("select count(*) from user_club where club_id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(clubID)
	err = row.Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	return
}
