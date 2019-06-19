package club

import (
	"gconst"
	"lobbyserver/lobby"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// onSetName 更新俱乐部的名称
func onSetClubMemberRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")
	log.Println("onSetClubMemberRole, userID:", userID)

	var query = r.URL.Query()
	clubID := query.Get("clubID")
	memberID := query.Get("memberID")
	role := query.Get("role") // 新俱乐部名称

	if clubID == "" {
		log.Println("onSetClubMemberRole, need club id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if role == "" {
		log.Println("onSetClubMemberRole, need role")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if memberID == "" {
		log.Println("onSetClubMemberRole, need member id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	roleInt, _ := strconv.Atoi(role)
	if roleInt != int(ClubRoleType_CRoleTypeMgr) && roleInt != int(ClubRoleType_CRoleTypeMember) {
		log.Printf("onSetClubMemberRole, role %d not club member and mgr", roleInt)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Println("onJoinClub, club not exist for clubID:", clubID)
		sendGenericError(w, ClubOperError_CERR_Club_Not_Exist)
		return
	}

	myRole := clubMgr.getClubMemberRole(userID, clubID)
	// 只有群主可以设管理员
	if myRole != int32(ClubRoleType_CRoleTypeCreator) {
		log.Printf("User %s not club %s creator, can change club member role", userID, clubID)
		sendGenericError(w, ClubOperError_CERR_Club_Only_Owner_And_Mgr_Can_Set)
		return
	}

	memberRole := clubMgr.getClubMemberRole(memberID, clubID)
	if memberRole == int32(ClubRoleType_CRoleTypeNone) {
		log.Printf("member %s no in club %s, can not change role", memberID, clubID)
		sendGenericError(w, ClubOperError_CERR_User_Not_In_Club)
		return
	}

	if memberRole == int32(roleInt) {
		log.Printf("member %s in club %s have been role %d, can't set repeat", memberID, clubID, roleInt)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	mySQLUtil := lobby.MySQLUtil()
	errCode := mySQLUtil.ChangeClubMemberRole(memberID, clubID, int32(roleInt))
	if errCode != 0 {
		log.Error("db change club member role error errCode:", errCode)
	}


	member := club.mm[memberID]
	member.Role = int32(roleInt)

	conn := lobby.Pool().Get()
	defer conn.Close()

	if int32(roleInt) == int32(ClubRoleType_CRoleTypeMgr) {
		conn.Do("SADD", gconst.LobbyClubManager+clubID, memberID)
	}else {
		conn.Do("SREM", gconst.LobbyClubManager+clubID, memberID)
	}

	userIDs := []string{memberID}
	notifyType := int32(ClubNotifyType_CNotify_Change_Member_Role)
	clubNotify := &MsgClubNotify{}
	clubNotify.NotifyType = &notifyType
	clubNotify.ClubID = &clubID
	// 发送通知
	sendClubNotify(userIDs, clubNotify)

	// 操作成功
	sendGenericError(w, ClubOperError_CERR_OK)
}
