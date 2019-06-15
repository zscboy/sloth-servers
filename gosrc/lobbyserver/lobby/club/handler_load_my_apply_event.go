package club

import (
	"gconst"
	"lobbyserver/lobby"
	"net/http"
	"strconv"

	"github.com/garyburd/redigo/redis"
	proto "github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// onLoadEvents 加载事件
func onLoadMyApplyEvent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")

	log.Println("onLoadMyApplyEvent, userID:", userID)

	var query = r.URL.Query()
	// 俱乐部ID

	cursorStr := query.Get("cursor")
	cursor := 0
	if cursorStr != "" {
		cursor, _ = strconv.Atoi(cursorStr)
	}

	// 获得redis连接
	conn := lobby.Pool().Get()
	defer conn.Close()

	const maxLoad int = 10

	// 先加载ID
	vs, err := redis.Values(conn.Do("LRANGE", gconst.LobbyClubUserApplicantEventPrefix+userID, cursor,
		cursor+maxLoad-1))
	if err != nil && err != redis.ErrNil {
		log.Println("onLoadMyApplyEvent, redis error:", err)
		sendGenericError(w, ClubOperError_CERR_Database_IO)
		return
	}

	records := make([]*MsgClubApplyRecord, 0, maxLoad)

	for _, v := range vs {
		buf, _ := redis.Bytes(v, nil)
		applyRecord := &MsgClubApplyRecord{}
		err := proto.Unmarshal(buf, applyRecord)
		if err != nil {
			log.Error("Unmarshal MsgClubApplyRecord, err:", err)
			continue
		}

		records = append(records, applyRecord)

	}

	cursor32 := int32(cursor + len(vs))
	if len(vs) < maxLoad {
		cursor32 = 0
	}

	applyRecordReply := &MsgClubLoadApplyRecordReply{}
	applyRecordReply.Records = records
	applyRecordReply.Cursor = &cursor32

	b, err := proto.Marshal(applyRecordReply)
	if err != nil {
		log.Println("onLoadMyApplyEvent, marshal error:", err)
		sendGenericError(w, ClubOperError_CERR_Encode_Decode)
		return
	}

	sendMsgClubReply(w, ClubReplyCode_RCOperation, b)
}
