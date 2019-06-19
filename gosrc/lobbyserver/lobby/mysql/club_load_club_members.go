package mysql

import (
	"lobbyserver/lobby/club"
)

// 拉取牌友群的成员ID，user_club表，user_id和club_id都是索引，可以从user_id查club_id,也可以从club_id查user_id
func loadClubMembers(clubID string) map[string]*club.Member {
	stmt, err := dbConn.Prepare("select user_id, club_role, allow_create_room from user_club where club_id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	members := make(map[string]*club.Member)
	rows, err := stmt.Query(clubID)
	defer rows.Close()

	for rows.Next() {
		var userID string
		var clubRole int32
		var allowCreateRoome int32
		err = rows.Scan(&userID, &clubRole, &allowCreateRoome)
		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}

		member := &club.Member{}
		member.Role = clubRole

		if allowCreateRoome > 0 {
			member.IsAllowCreateRoom = true
		}

		members[userID] = member
	}

	return members
}
