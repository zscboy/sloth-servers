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
func onLoadMyClubs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")

	log.Println("onLoadMyClubs, userID:", userID)

	// TODO: 用户的牌友群是否可以放在redis中
	mySQLUtil := lobby.MySQLUtil()
	clubIDs := mySQLUtil.LoadUserClubIDs(userID)

	reply := &MsgClubLoadMyClubsReply{}

	if len(clubIDs) > 0 {
		log.Printf("user %s load clubs, count:%d\n", userID, len(clubIDs))

		msgClubInfos := make([]*MsgClubInfo, 0, len(clubIDs))
		// removedClubIds := make([]string, 0, len(clubIDs))
		// check and load club from redis if need
		for _, cid := range clubIDs {
			if cid == "" {
				continue
			}

			// 检查牌友群是否在列表中，否则从数据库加载
			club, ok := clubMgr.clubs[cid]
			if !ok {
				clubInfo := mySQLUtil.LoadClubInfo(cid)
				if clubInfo == nil {
					log.Panic("onLoadMyClubs, no club for clubID:", cid)
				}

				club = newBaseClub(clubInfo.(*MsgClubInfo), cid)
				clubMgr.clubs[cid] = club
			}

			// clubInfo := club.constructMsgClubInfo()
			msgClubInfos = append(msgClubInfos, club.clubInfo)
		}

		// if len(removedClubIds) > 0 {
		// 	conn.Send("MULTI")
		// 	// 表明已经有俱乐部不存在，因此需要删除掉
		// 	for _, v := range removedClubIds {
		// 		conn.Send("SREM", stateless.PlayerClubSetPrefix+userID, v)
		// 	}
		// 	conn.Do("EXEC")
		// }

		// 检查是否未读事件
		if len(msgClubInfos) > 0 {
			conn := lobby.Pool().Get()
			defer conn.Close()

			conn.Send("MULTI")

			for _, cinfo := range msgClubInfos {
				clubID := cinfo.BaseInfo.GetClubID()
				conn.Send("SCARD", gconst.LobbyClubUnReadEventUserSetPrefix+clubID+":"+userID)
				conn.Send("SMEMBERS", gconst.LobbyClubManager+clubID)
				conn.Send("SCARD", gconst.LobbyClubMemberSetPrefix+clubID)
			}

			vs, err := redis.Values(conn.Do("EXEC"))

			if err == nil {
				for i, cinfo := range msgClubInfos {
					countEvent, _ := redis.Int(vs[i*2], nil)
					hasUnReadEvents := false
					if countEvent > 0 {
						hasUnReadEvents = true
					}
					cinfo.HasUnReadEvents = &hasUnReadEvents
					cinfo.Managers, _ = redis.Strings(vs[i*2+1], nil)
					memberCount, _ :=  redis.Int(vs[i*2+2], nil)
					memberCountInt32 := int32(memberCount)
					cinfo.MemberCount = &memberCountInt32
				}
			}
		}

		reply.Clubs = msgClubInfos
	}

	log.Println("reply:", reply)

	b, err := proto.Marshal(reply)
	if err != nil {
		log.Println("onCreateClub, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}
