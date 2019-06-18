package club

import (
	"gconst"
	"lobbyserver/lobby"
	"net/http"

	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// onLoadMyClubs 加载自己的俱乐部
func onLoadClub(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")

	log.Println("onLoadMyClubs, userID:", userID)

	var query = r.URL.Query()
	// 俱乐部ID
	clubID := query.Get("clubID")

	if clubID == "" {
		log.Println("onLoadClub, need club id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	mySQLUtil := lobby.MySQLUtil()

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		clubInfo := mySQLUtil.LoadClubInfo(clubID)
		if clubInfo == nil {
			sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
			return
		}

		club = newBaseClub(clubInfo.(*MsgClubInfo), clubID)
		clubMgr.clubs[clubID] = club
	}

	clubInfo := club.clubInfo

	memberCount := int32(mySQLUtil.CountClubUser(clubID))
	clubInfo.MemberCount = &memberCount

	conn := lobby.Pool().Get()
	defer conn.Close()

	conn.Send("MULTI")
	conn.Send("SCARD", gconst.LobbyClubUnReadEventUserSetPrefix+clubID+":"+userID)
	conn.Send("SMEMBERS", gconst.LobbyClubManager+clubID)
	vs, err := redis.Values(conn.Do("EXEC"))
	if err == nil {
		countEvent, _ := redis.Int(vs[0], nil)
		hasUnReadEvents := false
		if countEvent > 0 {
			hasUnReadEvents = true
		}
		clubInfo.HasUnReadEvents = &hasUnReadEvents
		clubInfo.Managers, _ = redis.Strings(vs[1], nil)
	}

	b, err := proto.Marshal(clubInfo)
	if err != nil {
		log.Println("onCreateClub, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}
