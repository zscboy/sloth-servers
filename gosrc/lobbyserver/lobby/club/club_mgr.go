package club

import (
	log "github.com/sirupsen/logrus"
)

// Member 牌友群成员
type Member struct {
	isOnline bool
	Role int32
	IsAllowCreateRoom bool
}
// Club 牌友群
type Club struct {
	ID       string
	clubInfo *MsgClubInfo    // club info
	mm       map[string]*Member // 成员列表
}

// MyClubMgr 用户管理
type MyClubMgr struct {
	clubs map[string]*Club
}

func newBaseClub(clubInfo *MsgClubInfo, clubID string) *Club {
	club := &Club{}
	club.clubInfo = clubInfo
	club.ID = clubID
	club.mm = make(map[string]*Member)

	return club
}

func newClubMgr() *MyClubMgr {
	clubMgr := &MyClubMgr{}
	clubMgr.clubs = make(map[string]*Club)
	return clubMgr
}


func (mgr *MyClubMgr) getClubMembeIDs(clubID string) []string {
	club, ok := clubMgr.clubs[clubID]
	if !ok {
		return make([]string, 0)
	}

	memberIDs := make([]string, 0, len(club.mm))

	for memberID := range club.mm {
		memberIDs = append(memberIDs, memberID)
	}

	return memberIDs
}

// GetClub 获取牌友圈成员的角色
func (mgr *MyClubMgr) getClubMemberRole(memberID string, clubID string) int32 {
	club, ok := clubMgr.clubs[clubID]
	if !ok {
		log.Printf("getClubMemberRole, club %s not exist", clubID)
		return int32(ClubRoleType_CRoleTypeNone)
	}

	member, ok := club.mm[memberID]
	if !ok {
		log.Printf("getClubMemberRole, member %s not in club %s", memberID, clubID)
		return int32(ClubRoleType_CRoleTypeNone)
	}

	return member.Role
}


// GetClub 获取牌友圈
func (mgr *MyClubMgr) GetClub(clubID string) interface{} {
	club, ok := clubMgr.clubs[clubID]
	if !ok {
		return nil
	}
	return club
}


// IsUserPermisionCreateRoom 判断用户是否有权限创建房间
func (mgr *MyClubMgr) IsUserPermisionCreateRoom(userID string, clubID string) bool {
	club, ok := clubMgr.clubs[clubID]
	if !ok {
		return false
	}

	member, ok := club.mm[userID]
	if !ok {
		return false
	}

	// if member.Role == int32(ClubRoleType_CRoleTypeCreator) || member.Role == int32(ClubRoleType_CRoleTypeMgr) {
	// 	return true
	// }

	if member.IsAllowCreateRoom {
		return true
	}

	return false
}

// IsUserPermisionDeleteRoom 判断用户是否有权限创建房间
func (mgr *MyClubMgr) IsUserPermisionDeleteRoom(userID string, clubID string) bool {
	role := mgr.getClubMemberRole(userID, clubID)
	if role == int32(ClubRoleType_CRoleTypeCreator) || role == int32(ClubRoleType_CRoleTypeMgr) {
		return true
	}

	return false

}

// IsClubMember 判断是否是牌友圈成员
func (mgr *MyClubMgr) IsClubMember(userID string, clubID string) bool {
	role := mgr.getClubMemberRole(userID, clubID)
	if role == int32(ClubRoleType_CRoleTypeNone) {
		return false
	}

	return true

}

// func (mgr *MyClubMgr)addClub(club *Club) {
// 	_, ok := mgr.clubs[club.ID]
// 	if !ok {
// 		mgr.clubs[club.ID] = club
// 	}
// }
