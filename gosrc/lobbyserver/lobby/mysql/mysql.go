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

func (*mySQLUtil) GetUserIDBy(account string) uint64 {
	return getUserIDBy(account)
}

func (*mySQLUtil) GetPasswordBy(account string) string {
	return getPasswordBy(account)
}

func (*mySQLUtil) GetOrGenerateUserID(account string) (userID uint64, isNew bool) {
	return getOrGenerateUserID(account)
}

func (*mySQLUtil) RegisterAccount(userID uint64, account string, passwd string, phone string, clientInfo *lobby.ClientInfo) error {
	return registerAccount(userID, account, passwd, phone, clientInfo)
}

func (*mySQLUtil) LoadUserInfo(userID uint64) *lobby.UserInfo {
	return loadUserInfo(userID)
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