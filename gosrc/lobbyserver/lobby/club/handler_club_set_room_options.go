package club

import (
	"log"
	"net/http"
	"gconst"
	"lobbyserver/lobby"
	proto "github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
)

// onSetRoomOptions 更新俱乐部的游戏房间设置
func onSetRoomOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")

	log.Println("onSetRoomOptions, userID:", userID)

	var query = r.URL.Query()
	clubID := query.Get("clubID")
	optionsStr := query.Get("options") // 创建房间的选项
	// payRoomOptionStr := query.Get("payOption")   // 支付选项

	if clubID == "" {
		log.Println("onSetRoomOptions, need club id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if optionsStr == "" {
		log.Println("onSetRoomOptions, need a valid options")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Printf("onSetRoomOptions, club %s not found", clubID)
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	clubInfo := club.clubInfo
	clubInfo.CreateRoomOptions = &optionsStr

	conn := lobby.Pool().Get()
	defer conn.Close()
	conn.Do("HSET", gconst.LobbyClubconfig + clubID, "options", optionsStr)

	b, err := proto.Marshal(clubInfo)
	if err != nil {
		log.Println("onCreateClub, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}
