package club

import (
	"lobbyserver/lobby"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// onSetName 更新俱乐部的名称
func onSetName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")
	log.Println("onSetName, userID:", userID)

	var query = r.URL.Query()
	clubID := query.Get("clubID")
	clname := query.Get("clname") // 新俱乐部名称

	if clubID == "" {
		log.Println("onSetName, need club id")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if clname == "" {
		log.Println("onSetName, need club name")
		sendGenericError(w, ClubOperError_CERR_Invalid_Input_Parameter)
		return
	}

	if len(clname) > 20 {
		log.Println("onSetName, clname too long:", len(clname))
		sendGenericError(w, ClubOperError_CERR_Club_Name_Too_Long)
		return
	}

	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Println("onJoinClub, club not exist for clubID:", clubID)
		sendGenericError(w, ClubOperError_CERR_Club_Not_Exist)
		return
	}

	if club.clubInfo.GetBaseInfo().GetClubName() == clname {
		sendGenericError(w, ClubOperError_CERR_Club_Name_Exist)
		return
	}

	role := clubMgr.getClubMemberRole(userID, clubID)
	// 只有群主和管理员可以设置群
	if role != int32(ClubRoleType_CRoleTypeCreator) && role != int32(ClubRoleType_CRoleTypeMgr) {
		sendGenericError(w, ClubOperError_CERR_Club_Only_Owner_And_Mgr_Can_Set)
		return
	}

	mySQLUtil := lobby.MySQLUtil()
	errCode := mySQLUtil.RenameClub(clubID, clname)
	if errCode != 0 {
		log.Error("db rename club error, errCode:", errCode)
	}

	club.clubInfo.GetBaseInfo().ClubName = &clname

	// 操作成功
	sendGenericError(w, ClubOperError_CERR_OK)
}
