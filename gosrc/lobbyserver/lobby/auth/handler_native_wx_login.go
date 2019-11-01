package auth

import (
	"lobbyserver/lobby"
	"lobbyserver/wechat"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func weiXinUserInfo2UserInof(wxUserInfo *wechat.LoadUserInfoReply) *lobby.UserInfo {
	userInfo := &lobby.UserInfo{}
	userInfo.OpenID = &wxUserInfo.OpenID
	userInfo.NickName = &wxUserInfo.NickName
	gender := uint32(wxUserInfo.Gender)
	userInfo.Gender = &gender
	userInfo.HeadImgUrl = &wxUserInfo.HeadImgURL

	return userInfo
}

// 微信开放平台接入微信登录
func handlerNativeWxLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wxCode := r.URL.Query().Get("code")

	log.Println("handlerNativeWxLogin, wxCode:", wxCode)

	loginReply := &lobby.MsgLoginReply{}

	accessTokenReply, err := wechat.LoadNativeAccessTokenFromWeChatServer(wxCode)
	if err != nil {
		log.Panicln("handlerWxLogin loadAccessTokenFromWeChatServer err:", err)
		errCode := int32(lobby.LoginError_ErrWxAuthFailed)
		loginReply.Result = &errCode
		replyWxLogin(w, loginReply)

		return
	}

	// 检查结果
	if accessTokenReply.ErrorCode != 0 {
		log.Errorf("loadAccessTokenFromWeChatServer, wechat server reply error code:%d, msg:%s\n", accessTokenReply.ErrorCode, accessTokenReply.ErrorMsg)
		errCode := int32(lobby.LoginError_ErrWxAuthFailed)
		loginReply.Result = &errCode
		replyWxLogin(w, loginReply)

		return
	}

	userInfoReply, err := wechat.GetWeiXinUserInfoFromAccessToken(accessTokenReply.AccessToken, accessTokenReply.OpenID)
	if err != nil {
		log.Error("loadUserInfoFromWeChatServer err:", err)
		errCode := int32(lobby.LoginError_ErrDecodeUserInfoFailed)
		loginReply.Result = &errCode
		replyWxLogin(w, loginReply)

		return
	}

	userInfo := weiXinUserInfo2UserInof(userInfoReply)

	qMod := r.URL.Query().Get("qMod")
	modV := r.URL.Query().Get("modV")
	csVer := r.URL.Query().Get("csVer")
	lobbyVer := r.URL.Query().Get("lobbyVer")
	operatingSystem := r.URL.Query().Get("operatingSystem")
	operatingSystemFamily := r.URL.Query().Get("operatingSystemFamily")
	deviceUniqueIdentifier := r.URL.Query().Get("deviceUniqueIdentifier")
	deviceName := r.URL.Query().Get("deviceName")
	deviceModel := r.URL.Query().Get("deviceModel")
	network := r.URL.Query().Get("network")

	clientInfo := &lobby.ClientInfo{}
	clientInfo.QMod = &qMod
	clientInfo.ModV = &modV
	clientInfo.CsVer = &csVer
	clientInfo.LobbyVer = &lobbyVer
	clientInfo.OperatingSystem = &operatingSystem
	clientInfo.OperatingSystemFamily = &operatingSystemFamily
	clientInfo.DeviceUniqueIdentifier = &deviceUniqueIdentifier
	clientInfo.DeviceName = &deviceName
	clientInfo.DeviceModel = &deviceModel
	clientInfo.Network = &network

	mySQLUtil := lobby.MySQLUtil()
	userID, isNew := mySQLUtil.LoadOrGenerateUserID(userInfo.GetOpenID())
	userInfo.UserID = &userID

	if isNew {
		// TODO注册账号
		registerAccount(userInfo.GetOpenID(), "", userInfo, clientInfo)
	} else {
		// 更新用户信息
		updateWxUserInfo(userInfo, clientInfo)
	}

	roomUtil := lobby.RoomUtil()
	lastRoomInfo := roomUtil.LoadLastRoomInfo(userID)

	// 生成token给客户端
	tk := lobby.GenTK(userID)

	errCode := int32(lobby.LoginError_ErrLoginSuccess)

	loginReply.Result = &errCode
	loginReply.Token = &tk
	loginReply.UserInfo = userInfo
	loginReply.LastRoomInfo = lastRoomInfo
	replyWxLogin(w, loginReply)
}
