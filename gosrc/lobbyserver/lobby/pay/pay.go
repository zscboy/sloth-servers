package pay

import (
	"lobbyserver/lobby"
)

var (
	payUtil = &myPayUtil{}
)

// myPayUtil implements IPayUtil
type myPayUtil struct {
}

func (*myPayUtil) Refund2Users(roomID string, handFinish int, inGameUserIDs []string) bool {
	orders := refund2Users(roomID, handFinish, inGameUserIDs)
	return len(orders) > 0
}

func (*myPayUtil) DoPayAndSave2RedisWith(roomType int, roomConfigID string,
	roomID string, userID string, gameNo string) (remainDiamond int, errCode int32) {
	return doPayAndSave2RedisWith(roomType, roomConfigID, roomID, userID, gameNo)
}

func (*myPayUtil) Refund2UserAndSave2Redis(roomID string, userID string, handFinish int) {
	refund2UserAndSave2Redis(roomID, userID, handFinish)
}

// InitWith init
func InitWith() {
	lobby.SetPayUtil(payUtil)
}
