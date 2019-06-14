package club

import (
	"gconst"
	"lobbyserver/lobby"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// onKickOut 剔除某个成员
func onKickOut(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")
	log.Println("onKickOut, userID:", userID)

	var query = r.URL.Query()
	// 俱乐部ID
	clubID := query.Get("clubID")
	memberID := query.Get("memberID")

	if clubID == "" {
		log.Println("onKickOut, need clubID")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if memberID == "" {
		log.Println("onKickOut, need memberID")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Println("onJoinClub, club not exist for clubID:", clubID)
		sendGenericError(w, ClubOperError_CERR_Club_Not_Exist)
		return
	}

	conn := lobby.Pool().Get()
	defer conn.Close()

	if !isKickOutAble(clubID, conn, userID, w, memberID) {
		return
	}

	mySQLUtil := lobby.MySQLUtil()
	errCode := mySQLUtil.RemoveUserFromClub(memberID, clubID)
	if errCode != 0 {
		log.Error("onKickOut, RemoveUserFromClub error, errCode:", errCode)
		return
	}

	// 清理玩家关于本club的数据
	redisClearClubUserData(clubID, memberID, conn)

	nick, _ := redis.String(conn.Do("HGET", gconst.LobbyUserTablePrefix+userID, "nick"))
	if nick == "" {
		nick = userID
	}

	clname := club.clubInfo.GetBaseInfo().GetClubName()
	userIDs := []string{
		memberID,
	}
	// 给前俱乐部成员memberID发送一个邮件
	var text = "您被 " + nick + " 踢出 " + clname + " 俱乐部!"
	sendClubEventMails(userIDs, text)

	// 生成事件
	newKickoutEvent(clubID, memberID, conn)

	// 操作成功
	sendGenericError(w, ClubOperError_CERR_OK)
}

func isKickOutAble(clubID string, conn redis.Conn, userID string, w http.ResponseWriter, memberID string) bool {
	if userID == memberID {
		log.Printf("onKickOut, userID is equal to memberID which intend to be kickout\n")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return false
	}

	mySQLUtil := lobby.MySQLUtil()
	myRole := mySQLUtil.LoadUserClubRole(userID, clubID)
	// 只有群主和管理员才可以踢人
	if myRole != int32(ClubRoleType_CRoleTypeCreator) && myRole != int32(ClubRoleType_CRoleTypeMgr) {
		log.Printf("onKickOut, club:%s, userID:%s, only creator can kickout member\n", clubID, userID)
		sendGenericError(w, ClubOperError_CERR_Only_Creator_And_Mgr_Can_KickOut)
		return false
	}

	memberRole := mySQLUtil.LoadUserClubRole(memberID, clubID)
	if memberRole == int32(ClubRoleType_CRoleTypeNone) {
		sendGenericError(w, ClubOperError_CERR_User_Not_In_Club)
		return false
	}

	// 被踢的人只能是普通成员
	if memberRole != int32(ClubRoleType_CRoleTypeMember) {
		sendGenericError(w, ClubOperError_CERR_Only_Creator_And_Mgr_Can_KickOut)
		return false
	}

	return true
}

func newKickoutEvent(clubID string, userID string, conn redis.Conn) {
	// 事件ID
	cn, err := redis.Int64(conn.Do("HINCRBY", gconst.LobbyClubSysTable, "clubEventID", 1))
	if err != nil {
		log.Panicln("newJoinEvent alloc eventID failed, redis err:", err)
	}

	clubEvent := &MsgClubEvent{}
	evtType32 := int32(ClubEventType_CEVT_Kickout)
	clubEvent.EvtType = &evtType32
	to := ""
	clubEvent.To = &to
	generatedTime32 := uint32(time.Since(time2010).Seconds())
	clubEvent.GeneratedTime = &generatedTime32
	needHandle := false // 通知事件是需要处理的
	clubEvent.NeedHandle = &needHandle

	eventID32 := uint32(cn % int64(0x0ffffffff))
	clubEvent.Id = &eventID32

	clubEvent.UserID1 = &userID

	clubEventBytes, err := proto.Marshal(clubEvent)
	if err != nil {
		log.Panic(err)
	}

	// TODO: 后面需要增加裁剪如下各个列表的定时器

	conn.Send("MULTI")
	// 加入到俱乐部的信息列表
	conn.Send("HSET", gconst.LobbyClubEventTablePrefix+clubID, eventID32, clubEventBytes)
	conn.Send("LPUSH", gconst.LobbyClubEventListPrefix+clubID, eventID32)
	_, err = conn.Do("EXEC")
	if err != nil {
		log.Panic(err)
	}

	// 加入到成员的未读信息列表
	// KEYS[1] prefix
	// KEYS[2] member-set key
	// KEYS[3] eventID
	_, err = luaScriptInsertNewEvent.Do(conn, gconst.LobbyClubUnReadEventUserListPrefix+clubID+":",
		gconst.LobbyClubUnReadEventUserSetPrefix+clubID+":", gconst.LobbyClubMemberSetPrefix+clubID, eventID32)

	if err != nil {
		log.Panic(err)
	}
}
