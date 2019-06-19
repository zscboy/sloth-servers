package club

import (
	"lobbyserver/lobby"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
	"gconst"
	"strconv"
)

func replyMemberInfo(memberID string, club *Club, w http.ResponseWriter,) {
	conn := lobby.Pool().Get()
	defer conn.Close()

	strValues, err := redis.Strings(conn.Do("HMGET", gconst.LobbyUserTablePrefix+memberID, "Nick", "Gender", "Protrait", "AvatarID"))
	if err != nil {
		log.Errorf("Redis get member %s info error %v", memberID, err)
		return
	}

	displayInfo := &MsgClubDisplayInfo{}
	nick := strValues[0]
	displayInfo.Nick = &nick
	gender, _ := strconv.Atoi(strValues[1])
	sex32 := uint32(gender)
	displayInfo.Gender = &sex32
	headIconURL := strValues[2]
	displayInfo.HeadIconURL = &headIconURL
	avatarID, _ := strconv.Atoi(strValues[3])
	avatarID32 := int32(avatarID)
	displayInfo.AvatarID = &avatarID32

	memberInfo := &MsgClubMemberInfo{}
	memberInfo.UserID = &memberID
	memberInfo.DisplayInfo = displayInfo

	role := club.mm[memberID].Role
	isAllowCreateRoom := club.mm[memberID].IsAllowCreateRoom
	memberInfo.Role = &role
	memberInfo.AllowCreateRoom = &isAllowCreateRoom

	b, err := proto.Marshal(memberInfo)
	if err != nil {
		log.Println("onLoadClubMembers, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}

// onSetName 更新俱乐部的名称
func onAllowMemberCreateRoom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")
	log.Println("onAllowMemberCreateRoom, userID:", userID)

	var query = r.URL.Query()
	clubID := query.Get("clubID")
	memberID := query.Get("memberID")
	allowCreateRoom := query.Get("allowCreateRoom") // 新俱乐部名称

	if clubID == "" {
		log.Println("onAllowMemberCreateRoom, need club id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if memberID == "" {
		log.Println("onAllowMemberCreateRoom, need member id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if allowCreateRoom == "" {
		log.Println("onAllowMemberCreateRoom, need allowCreateRoom")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	isAllowCreateRoom := false
	if allowCreateRoom == "yes" {
		isAllowCreateRoom = true
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Println("onAllowMemberCreateRoom, club not exist for clubID:", clubID)
		sendGenericError(w, ClubOperError_CERR_Club_Not_Exist)
		return
	}

	member, ok := club.mm[memberID]
	if !ok {
		log.Printf("onAllowMemberCreateRoom member %s not in club %s, can't change new room permision", memberID, clubID)
		sendGenericError(w, ClubOperError_CERR_User_Not_In_Club)
		return
	}

	if member.IsAllowCreateRoom == isAllowCreateRoom {
		log.Printf("onAllowMemberCreateRoom member %s in club %s have been permission %v, can't set repeat", memberID, clubID, isAllowCreateRoom)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	myRole := clubMgr.getClubMemberRole(userID, clubID)
	// 只有群主可以设管理员
	if myRole != int32(ClubRoleType_CRoleTypeCreator) && myRole != int32(ClubRoleType_CRoleTypeMgr) {
		log.Printf("onAllowMemberCreateRoom User %s not club %s creator and mgr, can change club member new room permission", userID, clubID)
		sendGenericError(w, ClubOperError_CERR_Club_Only_Owner_And_Mgr_Can_Set)
		return
	}

	permissionInt := 0
	if isAllowCreateRoom {
		permissionInt = 1
	}

	mySQLUtil := lobby.MySQLUtil()
	errCode := mySQLUtil.AllowMemberCreateRoom(memberID, clubID, int32(permissionInt))
	if errCode != 0 {
		log.Error("onAllowMemberCreateRoom db change club member role error errCode:", errCode)
	}

	member.IsAllowCreateRoom = isAllowCreateRoom
	log.Printf("isAllowCreateRoom:%v", isAllowCreateRoom)

	userIDs := []string{memberID}
	notifyType := int32(ClubNotifyType_CNotify_Change_Member_Role)
	clubNotify := &MsgClubNotify{}
	clubNotify.NotifyType = &notifyType
	clubNotify.ClubID = &clubID
	// 发送通知
	sendClubNotify(userIDs, clubNotify)

	// 操作成功
	// sendGenericError(w, ClubOperError_CERR_OK)

	replyMemberInfo(memberID, club, w)
}
