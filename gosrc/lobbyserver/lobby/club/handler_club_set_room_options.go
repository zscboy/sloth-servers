package club

import (
	"log"
	"net/http"
	"crypto/md5"
	"fmt"
	"gconst"
	"lobbyserver/lobby"
	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
)

// onSetRoomOptions 更新俱乐部的游戏房间设置
func onSetRoomOptions(w http.ResponseWriter, r *http.Request, userID string) {
	log.Println("onSetRoomOptions, userID:", userID)

	var query = r.URL.Query()
	clubID := query.Get("clubID")
	roomConfigStr := query.Get("roomConfig") // 创建房间的选项
	// payRoomOptionStr := query.Get("payOption")   // 支付选项

	if clubID == "" {
		log.Println("onSetRoomOptions, need club id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if roomConfigStr == "" {
		log.Println("onSetRoomOptions, need a valid croption")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Printf("onJoinApprove, club %s not found", clubID)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	bytes := []byte(roomConfigStr)
	md5Value := md5.Sum(bytes)
	roomConfigID := fmt.Sprintf("%x", md5Value)

	// 获得redis连接
	conn := lobby.Pool().Get()
	defer conn.Close()

	result, _ := redis.Int(conn.Do("HEXISTS", gconst.LobbyRoomConfigTable, roomConfigID))
	if result != 1 {
		_, err := conn.Do("HSET", gconst.LobbyRoomConfigTable, roomConfigID, bytes)
		if err != nil {
			log.Println("save club room config err:", err)
			return
		}
	}

	_, exist := lobby.RoomConfigs[roomConfigID]
	if !exist {
		lobby.RoomConfigs[roomConfigID] = roomConfigStr
	}

	conn.Do("HSET", gconst.LobbyClubconfig + clubID, "roomConfigID", roomConfigID)


	clubInfo := club.clubInfo

	b, err := proto.Marshal(clubInfo)
	if err != nil {
		log.Println("onCreateClub, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}
