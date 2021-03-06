package mysql

import (
	"database/sql"
	"lobbyserver/lobby"

	log "github.com/sirupsen/logrus"
)

var (
	sqlUtil = &mySQLUtil{}
	dbConn  *sql.DB
)

// myRoomUtil implements IRoomUtil
type mySQLUtil struct {
}

func (*mySQLUtil) UpdateWxUserInfo(userInfo *lobby.UserInfo, clientInfo *lobby.ClientInfo) error {
	return updateWxUserInfo(userInfo, clientInfo)
}

func (*mySQLUtil) UpdateAccountUserInfo(account string, clientInfo *lobby.ClientInfo) error {
	return updateAccountUserInfo(account, clientInfo)
}

func (*mySQLUtil) LoadUserIDByAccount(account string) string {
	return loadUserIDByAccount(account)
}

func (*mySQLUtil) LoadPasswordByAccount(account string) string {
	return loadPasswordByAccount(account)
}

func (*mySQLUtil) LoadOrGenerateUserID(account string) (userID string, isNew bool) {
	return loadOrGenerateUserID(account)
}

func (*mySQLUtil) RegisterAccount(account string, passwd string, userInfo *lobby.UserInfo, clientInfo *lobby.ClientInfo) error {
	return registerAccount(account, passwd, userInfo, clientInfo)
}

func (*mySQLUtil) LoadUserInfo(userID string) *lobby.UserInfo {
	return loadUserInfo(userID)
}

func (*mySQLUtil) PayForRoom(userID string, pay int, roomID string) (errCode int, lastNum int64, orderID string) {
	return payForRoom(userID, pay, roomID)
}

func (*mySQLUtil) RefundForRoom(userID string, refund int, orderID string) (errCode int, lastNum int64) {
	return refundForRoom(userID, refund, orderID)
}

func (*mySQLUtil) UpdateDiamond(userID string, change int64) (lastNum int64, errCode int32) {
	return updateDiamond(userID, change)
}

func (*mySQLUtil) CountUserClubNumber(userID string) (count int) {
	return countUserClub(userID)
}

func (*mySQLUtil) CreateClub(clubName string, creator string, isLeague int, wanka int, candy int, maxMember int) (clubID string, clubNumber string, errCode int) {
	return createClub(clubName, creator, isLeague, wanka, candy, maxMember)
}

func (*mySQLUtil) LoadClubUserIDs(clubID string) (userIDs []string) {
	return loadClubUserIDs(clubID)
}

func (*mySQLUtil) LoadUserClubIDs(userID string) (clubIDs []string) {
	return loadUserClubIDs(userID)
}

func (*mySQLUtil) LoadClubInfo(clubID string) (clubInfo interface{}) {
	return loadClubInfo(clubID)
}

func (*mySQLUtil) LoadUserClubRole(userID string, clubID string) (role int32) {
	return loadUserClubRole(userID, clubID)
}

func (*mySQLUtil) DeleteClub(clubID string) (errCode int32) {
	return deleteClub(clubID)
}

func (*mySQLUtil) LoadClubIDByNumber(number string) string {
	return loadClubIDByNumber(number)
}

func (*mySQLUtil) AddUserToClub(userID string, clubID string, role int32) (errCode int) {
	return addUserToClub(userID, clubID, role)
}

func (*mySQLUtil) LoadClubInfos(cursor int, count int) (clubInfos interface{}) {
	return loadClubInfos(cursor, count)
}

func (*mySQLUtil) RemoveUserFromClub(userID string, clubID string) (errCode int) {
	return removeUserFromClub(userID, clubID)
}

func (*mySQLUtil) LoadUserDiamond(userID string) int64 {
	return loadUserDiamond(userID)
}

func (*mySQLUtil) CountClubUser(clubID string) int {
	return countClubUser(clubID)
}

// errCode = 1 参数错误, 0 成功
func (*mySQLUtil) RenameClub(clubID string, clubName string) (errCode int) {
	return renameClub(clubID, clubName)
}

// errCode = 1 参数错误， 0 成功
func (*mySQLUtil) ChangeClubMemberRole(memberID string, clubID string, role int32) (errCode int) {
	return changeClubMemberRole(memberID, clubID, role)
}

func (*mySQLUtil) LoadClubMembers(clubID string) interface{} {
	return loadClubMembers(clubID)
}

// errCode = 1 参数错误, 0 成功
func (*mySQLUtil) AllowMemberCreateRoom(memberID string, clubID string, permission int32)  (errCode int) {
	return allowMemberCreateRoom(memberID, clubID, permission)
}
// InitWith init
func InitWith() {
	lobby.SetMySQLUtil(sqlUtil)

	conn, err := startMySQL()
	if err != nil {
		// log.Panic("StartMssql error ", err)
		log.Warn("StartMssql error ", err)
	}
	dbConn = conn

}
