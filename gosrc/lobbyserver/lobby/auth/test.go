package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"

	// "fmt"
	"lobbyserver/lobby"
	// "github.com/golang/protobuf/proto"
)

func handlerTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// log.Println("handlerTest")
	// qMod := "loginMode"
	// modV := "1.0.1"
	// csVer := "1.0.0"
	// lobbyVer := "1.0.0"
	// operatingSystem := "IOS"
	// operatingSystemFamily := "iphone5"
	// deviceUniqueIdentifier := "111222"
	// deviceName := "chen_phone"
	// deviceModel := "phone"
	// network := "4G"

	// clientInfo := &lobby.ClientInfo{}
	// clientInfo.QMod = &qMod
	// clientInfo.ModV = &modV
	// clientInfo.CsVer = &csVer
	// clientInfo.LobbyVer = &lobbyVer
	// clientInfo.OperatingSystem = &operatingSystem
	// clientInfo.OperatingSystemFamily = &operatingSystemFamily
	// clientInfo.DeviceUniqueIdentifier = &deviceUniqueIdentifier
	// clientInfo.DeviceName = &deviceName
	// clientInfo.DeviceModel = &deviceModel
	// clientInfo.Network = &network

	// userInfo := &lobby.WxUserInfo{}

	// // 保存用户信息
	// saveUserInfo(userInfo, clientInfo)

	// w.Write([]byte("ok"))
	mySQLUtil := lobby.MySQLUtil()
	userID, isNew := mySQLUtil.LoadOrGenerateUserID("12333")
	log.Println("userID:", userID, isNew)
}
