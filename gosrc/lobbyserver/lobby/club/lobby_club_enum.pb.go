// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lobby_club_enum.proto

package club

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClubReplyCode int32

const (
	ClubReplyCode_RCNone      ClubReplyCode = 0
	ClubReplyCode_RCError     ClubReplyCode = 1
	ClubReplyCode_RCOperation ClubReplyCode = 2
)

var ClubReplyCode_name = map[int32]string{
	0: "RCNone",
	1: "RCError",
	2: "RCOperation",
}

var ClubReplyCode_value = map[string]int32{
	"RCNone":      0,
	"RCError":     1,
	"RCOperation": 2,
}

func (x ClubReplyCode) Enum() *ClubReplyCode {
	p := new(ClubReplyCode)
	*p = x
	return p
}

func (x ClubReplyCode) String() string {
	return proto.EnumName(ClubReplyCode_name, int32(x))
}

func (x *ClubReplyCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubReplyCode_value, data, "ClubReplyCode")
	if err != nil {
		return err
	}
	*x = ClubReplyCode(value)
	return nil
}

func (ClubReplyCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953a149cff643cc3, []int{0}
}

// 俱乐部角色定义
type ClubRoleType int32

const (
	ClubRoleType_CRoleTypeNone    ClubRoleType = 0
	ClubRoleType_CRoleTypeMember  ClubRoleType = 1
	ClubRoleType_CRoleTypeCreator ClubRoleType = 2
	// CRoleTypeOwner = 4; // 拥有者
	ClubRoleType_CRoleTypeMgr ClubRoleType = 3
)

var ClubRoleType_name = map[int32]string{
	0: "CRoleTypeNone",
	1: "CRoleTypeMember",
	2: "CRoleTypeCreator",
	3: "CRoleTypeMgr",
}

var ClubRoleType_value = map[string]int32{
	"CRoleTypeNone":    0,
	"CRoleTypeMember":  1,
	"CRoleTypeCreator": 2,
	"CRoleTypeMgr":     3,
}

func (x ClubRoleType) Enum() *ClubRoleType {
	p := new(ClubRoleType)
	*p = x
	return p
}

func (x ClubRoleType) String() string {
	return proto.EnumName(ClubRoleType_name, int32(x))
}

func (x *ClubRoleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubRoleType_value, data, "ClubRoleType")
	if err != nil {
		return err
	}
	*x = ClubRoleType(value)
	return nil
}

func (ClubRoleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953a149cff643cc3, []int{1}
}

type ClubOperError int32

const (
	ClubOperError_CERR_OK                               ClubOperError = 0
	ClubOperError_CERR_Exceed_Max_Club_Count_Limit      ClubOperError = 1
	ClubOperError_CERR_No_Valid_Club_Number             ClubOperError = 2
	ClubOperError_CERR_Database_IO                      ClubOperError = 3
	ClubOperError_CERR_Encode_Decode                    ClubOperError = 4
	ClubOperError_CERR_Invalid_Input_Parameter          ClubOperError = 5
	ClubOperError_CERR_Only_Creator_And_Mgr_Can_KickOut ClubOperError = 6
	ClubOperError_CERR_You_Already_In_Club              ClubOperError = 7
	ClubOperError_CERR_You_Are_In_Club_Block_List       ClubOperError = 8
	ClubOperError_CERR_You_Already_Applicate            ClubOperError = 9
	ClubOperError_CERR_Invitee_Already_In_Club          ClubOperError = 10
	ClubOperError_CERR_Invitee_Are_In_Club_Block_List   ClubOperError = 11
	ClubOperError_CERR_Invitee_Already_Applicate        ClubOperError = 12
	ClubOperError_CERR_Club_Not_Exist                   ClubOperError = 13
	ClubOperError_CERR_Only_Creator_Can_Invite          ClubOperError = 14
	ClubOperError_CERR_Only_Creator_And_Mgr_Can_Approve ClubOperError = 15
	ClubOperError_CERR_No_Applicant                     ClubOperError = 16
	ClubOperError_CERR_Applicant_Already_In_Club        ClubOperError = 17
	ClubOperError_CERR_Applicant_In_Club_Block_List     ClubOperError = 18
	ClubOperError_CERR_Token_Invalid                    ClubOperError = 19
	ClubOperError_CERR_Club_Name_Too_Long               ClubOperError = 20
	ClubOperError_CERR_Club_Name_Exist                  ClubOperError = 21
	ClubOperError_CERR_Club_Only_Owner_Can_Disband      ClubOperError = 22
	ClubOperError_CERR_Owner_Can_not_quit               ClubOperError = 23
	ClubOperError_CERR_User_Not_In_Club                 ClubOperError = 24
	ClubOperError_CERR_Club_Only_Owner_And_Mgr_Can_Set  ClubOperError = 25
	ClubOperError_CERR_Club_Forbit_Join                 ClubOperError = 26
	ClubOperError_CERR_Input_Text_Too_Long              ClubOperError = 27
	ClubOperError_CERR_Club_Has_Room_In_PlayingState    ClubOperError = 28
	ClubOperError_CERR_Can_Not_Kick_Out_Creator_Or_Mgr  ClubOperError = 29
)

var ClubOperError_name = map[int32]string{
	0:  "CERR_OK",
	1:  "CERR_Exceed_Max_Club_Count_Limit",
	2:  "CERR_No_Valid_Club_Number",
	3:  "CERR_Database_IO",
	4:  "CERR_Encode_Decode",
	5:  "CERR_Invalid_Input_Parameter",
	6:  "CERR_Only_Creator_And_Mgr_Can_KickOut",
	7:  "CERR_You_Already_In_Club",
	8:  "CERR_You_Are_In_Club_Block_List",
	9:  "CERR_You_Already_Applicate",
	10: "CERR_Invitee_Already_In_Club",
	11: "CERR_Invitee_Are_In_Club_Block_List",
	12: "CERR_Invitee_Already_Applicate",
	13: "CERR_Club_Not_Exist",
	14: "CERR_Only_Creator_Can_Invite",
	15: "CERR_Only_Creator_And_Mgr_Can_Approve",
	16: "CERR_No_Applicant",
	17: "CERR_Applicant_Already_In_Club",
	18: "CERR_Applicant_In_Club_Block_List",
	19: "CERR_Token_Invalid",
	20: "CERR_Club_Name_Too_Long",
	21: "CERR_Club_Name_Exist",
	22: "CERR_Club_Only_Owner_Can_Disband",
	23: "CERR_Owner_Can_not_quit",
	24: "CERR_User_Not_In_Club",
	25: "CERR_Club_Only_Owner_And_Mgr_Can_Set",
	26: "CERR_Club_Forbit_Join",
	27: "CERR_Input_Text_Too_Long",
	28: "CERR_Club_Has_Room_In_PlayingState",
	29: "CERR_Can_Not_Kick_Out_Creator_Or_Mgr",
}

var ClubOperError_value = map[string]int32{
	"CERR_OK":                               0,
	"CERR_Exceed_Max_Club_Count_Limit":      1,
	"CERR_No_Valid_Club_Number":             2,
	"CERR_Database_IO":                      3,
	"CERR_Encode_Decode":                    4,
	"CERR_Invalid_Input_Parameter":          5,
	"CERR_Only_Creator_And_Mgr_Can_KickOut": 6,
	"CERR_You_Already_In_Club":              7,
	"CERR_You_Are_In_Club_Block_List":       8,
	"CERR_You_Already_Applicate":            9,
	"CERR_Invitee_Already_In_Club":          10,
	"CERR_Invitee_Are_In_Club_Block_List":   11,
	"CERR_Invitee_Already_Applicate":        12,
	"CERR_Club_Not_Exist":                   13,
	"CERR_Only_Creator_Can_Invite":          14,
	"CERR_Only_Creator_And_Mgr_Can_Approve": 15,
	"CERR_No_Applicant":                     16,
	"CERR_Applicant_Already_In_Club":        17,
	"CERR_Applicant_In_Club_Block_List":     18,
	"CERR_Token_Invalid":                    19,
	"CERR_Club_Name_Too_Long":               20,
	"CERR_Club_Name_Exist":                  21,
	"CERR_Club_Only_Owner_Can_Disband":      22,
	"CERR_Owner_Can_not_quit":               23,
	"CERR_User_Not_In_Club":                 24,
	"CERR_Club_Only_Owner_And_Mgr_Can_Set":  25,
	"CERR_Club_Forbit_Join":                 26,
	"CERR_Input_Text_Too_Long":              27,
	"CERR_Club_Has_Room_In_PlayingState":    28,
	"CERR_Can_Not_Kick_Out_Creator_Or_Mgr":  29,
}

func (x ClubOperError) Enum() *ClubOperError {
	p := new(ClubOperError)
	*p = x
	return p
}

func (x ClubOperError) String() string {
	return proto.EnumName(ClubOperError_name, int32(x))
}

func (x *ClubOperError) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubOperError_value, data, "ClubOperError")
	if err != nil {
		return err
	}
	*x = ClubOperError(value)
	return nil
}

func (ClubOperError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953a149cff643cc3, []int{2}
}

type ClubEventType int32

const (
	ClubEventType_CEVT_None         ClubEventType = 0
	ClubEventType_CEVT_ClubDisband  ClubEventType = 1
	ClubEventType_CEVT_NewApplicant ClubEventType = 2
	ClubEventType_CEVT_Approval     ClubEventType = 3
	ClubEventType_CEVT_Deny         ClubEventType = 4
	ClubEventType_CEVT_Join         ClubEventType = 5
	ClubEventType_CEVT_Quit         ClubEventType = 6
	ClubEventType_CEVT_Kickout      ClubEventType = 7
)

var ClubEventType_name = map[int32]string{
	0: "CEVT_None",
	1: "CEVT_ClubDisband",
	2: "CEVT_NewApplicant",
	3: "CEVT_Approval",
	4: "CEVT_Deny",
	5: "CEVT_Join",
	6: "CEVT_Quit",
	7: "CEVT_Kickout",
}

var ClubEventType_value = map[string]int32{
	"CEVT_None":         0,
	"CEVT_ClubDisband":  1,
	"CEVT_NewApplicant": 2,
	"CEVT_Approval":     3,
	"CEVT_Deny":         4,
	"CEVT_Join":         5,
	"CEVT_Quit":         6,
	"CEVT_Kickout":      7,
}

func (x ClubEventType) Enum() *ClubEventType {
	p := new(ClubEventType)
	*p = x
	return p
}

func (x ClubEventType) String() string {
	return proto.EnumName(ClubEventType_name, int32(x))
}

func (x *ClubEventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubEventType_value, data, "ClubEventType")
	if err != nil {
		return err
	}
	*x = ClubEventType(value)
	return nil
}

func (ClubEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953a149cff643cc3, []int{3}
}

type ClubFundEventType int32

const (
	ClubFundEventType_CFET_None            ClubFundEventType = 0
	ClubFundEventType_CFET_Add_By_Shop     ClubFundEventType = 1
	ClubFundEventType_CFET_Award_By_System ClubFundEventType = 3
	ClubFundEventType_CFET_Gift_By_System  ClubFundEventType = 4
	ClubFundEventType_CFET_Reduce_By_Room  ClubFundEventType = 5
	ClubFundEventType_CFET_Add_By_Room     ClubFundEventType = 6
)

var ClubFundEventType_name = map[int32]string{
	0: "CFET_None",
	1: "CFET_Add_By_Shop",
	3: "CFET_Award_By_System",
	4: "CFET_Gift_By_System",
	5: "CFET_Reduce_By_Room",
	6: "CFET_Add_By_Room",
}

var ClubFundEventType_value = map[string]int32{
	"CFET_None":            0,
	"CFET_Add_By_Shop":     1,
	"CFET_Award_By_System": 3,
	"CFET_Gift_By_System":  4,
	"CFET_Reduce_By_Room":  5,
	"CFET_Add_By_Room":     6,
}

func (x ClubFundEventType) Enum() *ClubFundEventType {
	p := new(ClubFundEventType)
	*p = x
	return p
}

func (x ClubFundEventType) String() string {
	return proto.EnumName(ClubFundEventType_name, int32(x))
}

func (x *ClubFundEventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubFundEventType_value, data, "ClubFundEventType")
	if err != nil {
		return err
	}
	*x = ClubFundEventType(value)
	return nil
}

func (ClubFundEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953a149cff643cc3, []int{4}
}

type ClubNotifyType int32

const (
	ClubNotifyType_CNotify_None                     ClubNotifyType = 0
	ClubNotifyType_CNotify_Change_Member_Role       ClubNotifyType = 1
	ClubNotifyType_CNotify_Allow_Member_Create_Room ClubNotifyType = 2
	ClubNotifyType_CNotify_New_Member_Apply         ClubNotifyType = 3
	ClubNotifyType_CNotify_Member_Join_Approval     ClubNotifyType = 4
)

var ClubNotifyType_name = map[int32]string{
	0: "CNotify_None",
	1: "CNotify_Change_Member_Role",
	2: "CNotify_Allow_Member_Create_Room",
	3: "CNotify_New_Member_Apply",
	4: "CNotify_Member_Join_Approval",
}

var ClubNotifyType_value = map[string]int32{
	"CNotify_None":                     0,
	"CNotify_Change_Member_Role":       1,
	"CNotify_Allow_Member_Create_Room": 2,
	"CNotify_New_Member_Apply":         3,
	"CNotify_Member_Join_Approval":     4,
}

func (x ClubNotifyType) Enum() *ClubNotifyType {
	p := new(ClubNotifyType)
	*p = x
	return p
}

func (x ClubNotifyType) String() string {
	return proto.EnumName(ClubNotifyType_name, int32(x))
}

func (x *ClubNotifyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubNotifyType_value, data, "ClubNotifyType")
	if err != nil {
		return err
	}
	*x = ClubNotifyType(value)
	return nil
}

func (ClubNotifyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953a149cff643cc3, []int{5}
}

func init() {
	proto.RegisterEnum("club.ClubReplyCode", ClubReplyCode_name, ClubReplyCode_value)
	proto.RegisterEnum("club.ClubRoleType", ClubRoleType_name, ClubRoleType_value)
	proto.RegisterEnum("club.ClubOperError", ClubOperError_name, ClubOperError_value)
	proto.RegisterEnum("club.ClubEventType", ClubEventType_name, ClubEventType_value)
	proto.RegisterEnum("club.ClubFundEventType", ClubFundEventType_name, ClubFundEventType_value)
	proto.RegisterEnum("club.ClubNotifyType", ClubNotifyType_name, ClubNotifyType_value)
}

func init() { proto.RegisterFile("lobby_club_enum.proto", fileDescriptor_953a149cff643cc3) }

var fileDescriptor_953a149cff643cc3 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4b, 0x53, 0x1b, 0x47,
	0x10, 0xb6, 0x1e, 0x40, 0xdc, 0x80, 0x69, 0x06, 0x30, 0x60, 0x63, 0xc7, 0xf1, 0x23, 0x0f, 0x1d,
	0x72, 0xcf, 0x51, 0x5e, 0x44, 0x42, 0x0c, 0xc8, 0x11, 0x0a, 0x55, 0xb9, 0x64, 0x6a, 0xa4, 0x6d,
	0xcb, 0x5b, 0xac, 0x66, 0x36, 0xa3, 0x59, 0x60, 0xff, 0x4a, 0xaa, 0x72, 0xcf, 0x4f, 0xca, 0xcf,
	0x49, 0x4d, 0xaf, 0x76, 0x96, 0x87, 0xaa, 0x7c, 0xda, 0xea, 0xfe, 0x7a, 0xbf, 0xfe, 0xfa, 0x31,
	0x0d, 0x3b, 0xa9, 0x19, 0x8d, 0x0a, 0x39, 0x4e, 0xf3, 0x91, 0x24, 0x9d, 0x4f, 0x7f, 0xcc, 0xac,
	0x71, 0x46, 0xb4, 0xbd, 0xa3, 0xf3, 0x13, 0xac, 0x47, 0x69, 0x3e, 0x1a, 0x50, 0x96, 0x16, 0x91,
	0x89, 0x49, 0x00, 0x2c, 0x0f, 0xa2, 0x33, 0xa3, 0x09, 0x1f, 0x89, 0x55, 0x58, 0x19, 0x44, 0x3d,
	0x6b, 0x8d, 0xc5, 0x86, 0xd8, 0x80, 0xd5, 0x41, 0xd4, 0xcf, 0xc8, 0x2a, 0x97, 0x18, 0x8d, 0xcd,
	0xce, 0x9f, 0xb0, 0xc6, 0xbf, 0x9a, 0x94, 0x86, 0x45, 0x46, 0x62, 0x13, 0xd6, 0xa3, 0xca, 0x98,
	0x13, 0x6c, 0xc1, 0x46, 0x70, 0x9d, 0xd2, 0x74, 0x44, 0x9e, 0x68, 0x1b, 0x30, 0x38, 0x23, 0x4b,
	0xca, 0x19, 0x8b, 0x4d, 0x81, 0xb0, 0x56, 0x87, 0x4e, 0x2c, 0xb6, 0x3a, 0xff, 0xad, 0x94, 0xda,
	0x7c, 0x4e, 0x16, 0xe1, 0xf5, 0x44, 0xbd, 0xc1, 0x40, 0xf6, 0x3f, 0xe0, 0x23, 0xf1, 0x16, 0x5e,
	0xb1, 0xd1, 0xbb, 0x19, 0x13, 0xc5, 0xf2, 0x54, 0xdd, 0x48, 0x1f, 0x2d, 0x23, 0x93, 0x6b, 0x27,
	0x4f, 0x92, 0x69, 0xe2, 0xb0, 0x21, 0x5e, 0xc0, 0x3e, 0x47, 0x9d, 0x19, 0x79, 0xa1, 0xd2, 0x24,
	0x2e, 0x63, 0xce, 0x72, 0xd6, 0xd2, 0x64, 0x2d, 0x1e, 0x3e, 0x54, 0x4e, 0x8d, 0xd4, 0x8c, 0xe4,
	0x71, 0x1f, 0x5b, 0xe2, 0x29, 0x88, 0x92, 0x5a, 0x8f, 0x4d, 0x4c, 0xf2, 0x90, 0xfc, 0x07, 0xdb,
	0xe2, 0x15, 0x1c, 0xb0, 0xff, 0x58, 0x5f, 0x31, 0xd7, 0xb1, 0xce, 0x72, 0x27, 0x3f, 0x2a, 0xab,
	0xa6, 0xe4, 0xc8, 0xe2, 0x92, 0xf8, 0x01, 0xde, 0x95, 0x0a, 0x75, 0x5a, 0xc8, 0x79, 0x71, 0xb2,
	0xab, 0x63, 0x79, 0x3a, 0xb1, 0x32, 0x52, 0x5a, 0x7e, 0x48, 0xc6, 0x97, 0xfd, 0xdc, 0xe1, 0xb2,
	0x38, 0x80, 0x3d, 0x0e, 0xfd, 0xc3, 0xe4, 0xb2, 0x9b, 0x5a, 0x52, 0x71, 0x21, 0x8f, 0x35, 0xeb,
	0xc3, 0x15, 0xf1, 0x06, 0xbe, 0xae, 0x51, 0x4b, 0x15, 0x22, 0xdf, 0xa7, 0x66, 0x7c, 0x29, 0x4f,
	0x92, 0x99, 0xc3, 0xaf, 0xc4, 0x4b, 0x78, 0xf6, 0x80, 0xa2, 0x9b, 0x65, 0x69, 0x32, 0x56, 0x8e,
	0xf0, 0xf1, 0x6d, 0xbd, 0x89, 0x23, 0x7a, 0x90, 0x06, 0xc4, 0x77, 0xf0, 0xe6, 0x6e, 0xc4, 0xe2,
	0x54, 0xab, 0xe2, 0x35, 0xbc, 0x5c, 0x48, 0x55, 0xa7, 0x5b, 0x13, 0xbb, 0xb0, 0xc5, 0x31, 0x65,
	0x8b, 0x8d, 0x93, 0xbd, 0x1b, 0xff, 0xf3, 0x7a, 0xd0, 0x71, 0xa7, 0x2b, 0xbe, 0x1b, 0x25, 0x1b,
	0x3e, 0xf9, 0x72, 0xdf, 0xba, 0x59, 0x66, 0xcd, 0x15, 0xe1, 0x86, 0xd8, 0x81, 0xcd, 0x6a, 0xa2,
	0xf3, 0xe4, 0xda, 0x21, 0x06, 0x81, 0xc1, 0xf7, 0xa0, 0xda, 0x4d, 0xf1, 0x0e, 0xbe, 0xb9, 0x17,
	0xb3, 0xa0, 0x56, 0x11, 0xc6, 0x3f, 0x34, 0x97, 0xa4, 0xab, 0x61, 0xe3, 0x96, 0x78, 0x0e, 0xbb,
	0xb7, 0xea, 0x53, 0x53, 0x92, 0x43, 0x63, 0xe4, 0x89, 0xd1, 0x13, 0xdc, 0x16, 0x7b, 0xb0, 0x7d,
	0x0f, 0x2c, 0xab, 0xdf, 0x09, 0x8b, 0xca, 0x08, 0x17, 0xd8, 0xbf, 0xd6, 0x54, 0x96, 0x75, 0x98,
	0xcc, 0x46, 0x4a, 0xc7, 0xf8, 0x34, 0x90, 0xd7, 0x98, 0x36, 0x4e, 0xfe, 0x95, 0x27, 0x0e, 0x77,
	0xc5, 0x3e, 0xec, 0x30, 0xf8, 0xfb, 0x8c, 0x2c, 0x77, 0xb6, 0xaa, 0x69, 0x4f, 0x7c, 0x0f, 0x6f,
	0x17, 0xb2, 0xdf, 0x6e, 0xde, 0x39, 0x39, 0xdc, 0x0f, 0x24, 0x1c, 0x79, 0x64, 0xec, 0x28, 0x71,
	0xf2, 0x57, 0x93, 0x68, 0x7c, 0x16, 0x76, 0xb1, 0x5c, 0xe8, 0x21, 0xdd, 0xb8, 0xba, 0xb4, 0xe7,
	0xe2, 0x5b, 0x78, 0x5d, 0xff, 0xf8, 0x8b, 0x9a, 0xc9, 0x81, 0x31, 0x53, 0x2f, 0xe1, 0x63, 0xaa,
	0x8a, 0x44, 0x4f, 0xce, 0x9d, 0x9f, 0xff, 0x41, 0x2d, 0x45, 0x69, 0x16, 0xe9, 0x77, 0x5d, 0xf6,
	0x73, 0x17, 0x06, 0xda, 0xb7, 0x5e, 0x12, 0xbe, 0xe8, 0xfc, 0xd3, 0x28, 0x9f, 0x76, 0xef, 0x8a,
	0xb4, 0xe3, 0xe3, 0xb1, 0x0e, 0x8f, 0xa3, 0xde, 0xc5, 0x50, 0xce, 0x0f, 0x07, 0xbf, 0xcb, 0x8b,
	0x21, 0xa7, 0xac, 0x7a, 0xd4, 0x28, 0x47, 0xef, 0x83, 0xe8, 0xba, 0x1e, 0x7d, 0x93, 0x0f, 0x8f,
	0x77, 0x97, 0x3b, 0xa2, 0x52, 0x6c, 0x05, 0xba, 0x43, 0xd2, 0x05, 0xb6, 0x83, 0xc9, 0xe5, 0x2e,
	0x05, 0xf3, 0x37, 0xdf, 0xdd, 0x65, 0x3e, 0x3d, 0xde, 0xf4, 0x7a, 0x4d, 0xee, 0x70, 0xa5, 0xf3,
	0x77, 0x03, 0x36, 0x7d, 0xea, 0xa3, 0x5c, 0xc7, 0x77, 0x35, 0x1e, 0xf5, 0xee, 0x68, 0xf4, 0x66,
	0x37, 0x8e, 0xe5, 0xfb, 0x42, 0x9e, 0x7f, 0x36, 0x19, 0x36, 0x78, 0x0f, 0xd8, 0x7b, 0xad, 0x6c,
	0xe9, 0x2f, 0x66, 0x8e, 0xa6, 0xd8, 0xe2, 0xe7, 0xe1, 0x91, 0x9f, 0x93, 0x4f, 0xee, 0x16, 0xd0,
	0x0e, 0xc0, 0x80, 0xe2, 0x7c, 0x4c, 0x1e, 0xf2, 0x0d, 0xc6, 0xa5, 0xfb, 0x19, 0xd8, 0xbb, 0xdc,
	0xf9, 0xb7, 0x01, 0x4f, 0xbc, 0xb8, 0x33, 0xe3, 0x92, 0x4f, 0x05, 0x2b, 0xf3, 0x15, 0x94, 0x66,
	0x25, 0xce, 0x9f, 0x86, 0xb9, 0x27, 0xfa, 0xac, 0xf4, 0x84, 0x64, 0x79, 0x7f, 0xa5, 0xbf, 0xb1,
	0xd8, 0xe0, 0xa5, 0x9c, 0xe3, 0xdd, 0x34, 0x35, 0xd7, 0x15, 0xcc, 0xa3, 0xa2, 0x32, 0x55, 0x93,
	0xf7, 0xa2, 0xe2, 0xa5, 0x10, 0xe3, 0x7b, 0x5f, 0x60, 0x8b, 0x9f, 0xf5, 0x1c, 0x9d, 0x23, 0xbe,
	0xbf, 0xf5, 0x18, 0xda, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xff, 0xa1, 0xc9, 0x1f, 0x7b, 0x06,
	0x00, 0x00,
}
