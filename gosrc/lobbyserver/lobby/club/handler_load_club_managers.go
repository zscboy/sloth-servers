package club

import (
	"net/http"
	"gconst"
	"lobbyserver/lobby"
	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// // ClubUser 牌友群成员信息
// type ClubUser struct {
// 	UserID string
// 	Role int32
// }

// onLoadClubMembers 加载俱乐部成员列表
func onLoadClubManagers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")

	log.Println("onLoadClubMembers, userID:", userID)

	var query = r.URL.Query()
	// 俱乐部ID
	clubID := query.Get("clubID")
	// cursorStr := query.Get("cursor")

	if clubID == "" {
		log.Println("onLoadClubMembers, clubID is empty, userID:", userID)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Println("onLoadClubMembers, clubID is invalid:", clubID)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	mySQLUtil := lobby.MySQLUtil()
	role := mySQLUtil.LoadUserClubRole(userID, clubID)
	if role == int32(ClubRoleType_CRoleTypeNone) {
		sendGenericError(w, ClubOperError_CERR_User_Not_In_Club)
		return
	}

	// 获得redis连接
	conn := lobby.Pool().Get()
	defer conn.Close()

	memberIDs, _ := redis.Strings(conn.Do("SMEMBERS", gconst.LobbyClubManager + clubID))

	loadMemberReply := &MsgClubLoadMembersReply{}
	newCusor32 := int32(0)
	loadMemberReply.Cursor = &newCusor32

	// 填充member 列表
	loadMemberReply.Members = constructClubMemberList(memberIDs, conn, club)

	b, err := proto.Marshal(loadMemberReply)
	if err != nil {
		log.Println("onLoadClubMembers, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}
