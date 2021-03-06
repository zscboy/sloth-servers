package mysql

import (
	"database/sql"
	"lobbyserver/lobby"
)

// 查询用户信息
func loadUserInfo(userID string) *lobby.UserInfo {
	stmt, err := dbConn.Prepare("select open_id, phone, nick_name, gender, provice, city, country, head_img_url from user where user_id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var openID sql.NullString
	var phone sql.NullString
	var nickName sql.NullString
	var gender sql.NullInt64
	var provice sql.NullString
	var city sql.NullString
	var country sql.NullString
	var headImgURL sql.NullString

	row := stmt.QueryRow(userID)
	err = row.Scan(&openID, &phone, &nickName, &gender, &provice, &city, &country, &headImgURL)
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(err.Error())
	}

	userInfo := &lobby.UserInfo{}
	userInfo.UserID = &userID

	if openID.Valid {
		userInfo.OpenID = &openID.String
	}

	if phone.Valid {
		userInfo.Phone = &phone.String
	}

	if nickName.Valid {
		userInfo.NickName = &nickName.String
	}

	if provice.Valid {
		userInfo.Province = &provice.String
	}

	if city.Valid {
		userInfo.City = &city.String
	}

	if country.Valid {
		userInfo.Country = &country.String
	}

	if headImgURL.Valid {
		userInfo.HeadImgUrl = &headImgURL.String
	}

	if gender.Valid {
		uint32Sex := uint32(gender.Int64)
		userInfo.Gender = &uint32Sex
	}

	return userInfo
}
