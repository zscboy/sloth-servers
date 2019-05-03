package room

import (
	"lobbyserver/lobby"
)

var (
	roomUtil = &myRoomUtil{}
)

// myRoomUtil implements IRoomUtil
type myRoomUtil struct {
}

func (*myRoomUtil) LoadLastRoomInfo(userID string) *lobby.RoomInfo {
	return loadLastRoomInfo(userID)
}

func (*myRoomUtil) LoadUserLastEnterRoomID(userID string) string {
	return loadUserLastEnterRoomID(userID)
}

func (*myRoomUtil) DeleteRoomInfoFromRedis(roomID string, userID string) {
	deleteRoomInfoFromRedis(roomID, userID)
}

// InitWith init
func InitWith() {
	lobby.SetRoomUtil(roomUtil)
	lobby.RegHTTPHandle("GET", "/createRoom", handlerCreateRoom)
	lobby.RegHTTPHandle("GET", "/requestRoomInfo", handlerRequestRoomInfo)
	// lobby.AccUserIDHTTPHandlers["/loadLastRoomInfo"] = handlerLoadLastRoomInfo // 拉取用户最后所在的房间
	lobby.RegHTTPHandle("GET", "/deleteRoom", handlerDeleteRoom) // 删除房间
}
