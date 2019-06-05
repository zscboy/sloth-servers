// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gconst.proto

package gconst

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

type SSMsgType int32

const (
	SSMsgType_Request  SSMsgType = 1
	SSMsgType_Response SSMsgType = 2
	SSMsgType_Notify   SSMsgType = 3
)

var SSMsgType_name = map[int32]string{
	1: "Request",
	2: "Response",
	3: "Notify",
}

var SSMsgType_value = map[string]int32{
	"Request":  1,
	"Response": 2,
	"Notify":   3,
}

func (x SSMsgType) Enum() *SSMsgType {
	p := new(SSMsgType)
	*p = x
	return p
}

func (x SSMsgType) String() string {
	return proto.EnumName(SSMsgType_name, int32(x))
}

func (x *SSMsgType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSMsgType_value, data, "SSMsgType")
	if err != nil {
		return err
	}
	*x = SSMsgType(value)
	return nil
}

func (SSMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{0}
}

type SSMsgReqCode int32

const (
	SSMsgReqCode_CreateRoom       SSMsgReqCode = 1
	SSMsgReqCode_DeleteRoom       SSMsgReqCode = 2
	SSMsgReqCode_RoomStateNotify  SSMsgReqCode = 3
	SSMsgReqCode_AAExitRoomNotify SSMsgReqCode = 4
	SSMsgReqCode_AAEnterRoom      SSMsgReqCode = 5
	SSMsgReqCode_Donate           SSMsgReqCode = 6
	SSMsgReqCode_UpdateLocation   SSMsgReqCode = 7
	SSMsgReqCode_UpdatePropCfg    SSMsgReqCode = 8
	SSMsgReqCode_HandBeginNotify  SSMsgReqCode = 9
)

var SSMsgReqCode_name = map[int32]string{
	1: "CreateRoom",
	2: "DeleteRoom",
	3: "RoomStateNotify",
	4: "AAExitRoomNotify",
	5: "AAEnterRoom",
	6: "Donate",
	7: "UpdateLocation",
	8: "UpdatePropCfg",
	9: "HandBeginNotify",
}

var SSMsgReqCode_value = map[string]int32{
	"CreateRoom":       1,
	"DeleteRoom":       2,
	"RoomStateNotify":  3,
	"AAExitRoomNotify": 4,
	"AAEnterRoom":      5,
	"Donate":           6,
	"UpdateLocation":   7,
	"UpdatePropCfg":    8,
	"HandBeginNotify":  9,
}

func (x SSMsgReqCode) Enum() *SSMsgReqCode {
	p := new(SSMsgReqCode)
	*p = x
	return p
}

func (x SSMsgReqCode) String() string {
	return proto.EnumName(SSMsgReqCode_name, int32(x))
}

func (x *SSMsgReqCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSMsgReqCode_value, data, "SSMsgReqCode")
	if err != nil {
		return err
	}
	*x = SSMsgReqCode(value)
	return nil
}

func (SSMsgReqCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{1}
}

type SSMsgError int32

const (
	SSMsgError_ErrSuccess                       SSMsgError = 0
	SSMsgError_ErrDecode                        SSMsgError = 1
	SSMsgError_ErrEncode                        SSMsgError = 2
	SSMsgError_ErrRoomExist                     SSMsgError = 3
	SSMsgError_ErrNoRoomConfig                  SSMsgError = 4
	SSMsgError_ErrUnsupportRoomType             SSMsgError = 5
	SSMsgError_ErrDecodeRoomConfig              SSMsgError = 6
	SSMsgError_ErrRoomNotExist                  SSMsgError = 7
	SSMsgError_ErrTakeoffDiamondFailedNotEnough SSMsgError = 9
	SSMsgError_ErrTakeoffDiamondFailedIO        SSMsgError = 10
	SSMsgError_ErrTakeoffDiamondFailedRepeat    SSMsgError = 11
	SSMsgError_ErrRoomIsNoEmpty                 SSMsgError = 12
)

var SSMsgError_name = map[int32]string{
	0:  "ErrSuccess",
	1:  "ErrDecode",
	2:  "ErrEncode",
	3:  "ErrRoomExist",
	4:  "ErrNoRoomConfig",
	5:  "ErrUnsupportRoomType",
	6:  "ErrDecodeRoomConfig",
	7:  "ErrRoomNotExist",
	9:  "ErrTakeoffDiamondFailedNotEnough",
	10: "ErrTakeoffDiamondFailedIO",
	11: "ErrTakeoffDiamondFailedRepeat",
	12: "ErrRoomIsNoEmpty",
}

var SSMsgError_value = map[string]int32{
	"ErrSuccess":                       0,
	"ErrDecode":                        1,
	"ErrEncode":                        2,
	"ErrRoomExist":                     3,
	"ErrNoRoomConfig":                  4,
	"ErrUnsupportRoomType":             5,
	"ErrDecodeRoomConfig":              6,
	"ErrRoomNotExist":                  7,
	"ErrTakeoffDiamondFailedNotEnough": 9,
	"ErrTakeoffDiamondFailedIO":        10,
	"ErrTakeoffDiamondFailedRepeat":    11,
	"ErrRoomIsNoEmpty":                 12,
}

func (x SSMsgError) Enum() *SSMsgError {
	p := new(SSMsgError)
	*p = x
	return p
}

func (x SSMsgError) String() string {
	return proto.EnumName(SSMsgError_name, int32(x))
}

func (x *SSMsgError) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SSMsgError_value, data, "SSMsgError")
	if err != nil {
		return err
	}
	*x = SSMsgError(value)
	return nil
}

func (SSMsgError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{2}
}

type RoomType int32

const (
	// 大丰麻将
	RoomType_DafengMJ RoomType = 1
	// 贯蛋
	RoomType_GuanDang RoomType = 2
	// 东台
	RoomType_DongTaiMJ RoomType = 3
	// 盐城
	RoomType_YanChengMJ RoomType = 4
	// 韶关
	RoomType_ShaoGuanMJ RoomType = 5
	// 宁安
	RoomType_NingAnMJ RoomType = 6
	// 新疆杠后
	RoomType_XinJiangGH RoomType = 7
	// 大丰关张
	RoomType_DafengGZ RoomType = 8
	// 大丰7王523
	RoomType_Dafeng7w523 RoomType = 9
	// 牛牛
	RoomType_NiuNiu RoomType = 10
	// 斗地主
	RoomType_DDZ RoomType = 11
	// 新疆血流
	RoomType_XueLiuMJ RoomType = 12
	// 兰州麻将
	RoomType_LanZhouMJ RoomType = 13
	// 老兰州
	RoomType_LLanZouMJ RoomType = 14
	// 张掖
	RoomType_ZhangYeMJ RoomType = 15
	// 泰安
	RoomType_TacnMJ RoomType = 16
	// 泰安吹牛扑克牌
	RoomType_TacnPok RoomType = 17
	// 跑胡子
	RoomType_PHZ RoomType = 18
	// 法库麻将
	RoomType_FKMJ RoomType = 19
	// 法库麻将非混
	RoomType_FUMJ RoomType = 20
	// 湛江麻将
	RoomType_ZJMJ RoomType = 21
)

var RoomType_name = map[int32]string{
	1:  "DafengMJ",
	2:  "GuanDang",
	3:  "DongTaiMJ",
	4:  "YanChengMJ",
	5:  "ShaoGuanMJ",
	6:  "NingAnMJ",
	7:  "XinJiangGH",
	8:  "DafengGZ",
	9:  "Dafeng7w523",
	10: "NiuNiu",
	11: "DDZ",
	12: "XueLiuMJ",
	13: "LanZhouMJ",
	14: "LLanZouMJ",
	15: "ZhangYeMJ",
	16: "TacnMJ",
	17: "TacnPok",
	18: "PHZ",
	19: "FKMJ",
	20: "FUMJ",
	21: "ZJMJ",
}

var RoomType_value = map[string]int32{
	"DafengMJ":    1,
	"GuanDang":    2,
	"DongTaiMJ":   3,
	"YanChengMJ":  4,
	"ShaoGuanMJ":  5,
	"NingAnMJ":    6,
	"XinJiangGH":  7,
	"DafengGZ":    8,
	"Dafeng7w523": 9,
	"NiuNiu":      10,
	"DDZ":         11,
	"XueLiuMJ":    12,
	"LanZhouMJ":   13,
	"LLanZouMJ":   14,
	"ZhangYeMJ":   15,
	"TacnMJ":      16,
	"TacnPok":     17,
	"PHZ":         18,
	"FKMJ":        19,
	"FUMJ":        20,
	"ZJMJ":        21,
}

func (x RoomType) Enum() *RoomType {
	p := new(RoomType)
	*p = x
	return p
}

func (x RoomType) String() string {
	return proto.EnumName(RoomType_name, int32(x))
}

func (x *RoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RoomType_value, data, "RoomType")
	if err != nil {
		return err
	}
	*x = RoomType(value)
	return nil
}

func (RoomType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{3}
}

type RoomState int32

const (
	// RoomIdle 房间空闲
	RoomState_SRoomIdle RoomState = 0
	// RoomWaiting 房间正在等待玩家进入
	RoomState_SRoomWaiting RoomState = 1
	// RoomPlaying 游戏正在进行中
	RoomState_SRoomPlaying RoomState = 2
)

var RoomState_name = map[int32]string{
	0: "SRoomIdle",
	1: "SRoomWaiting",
	2: "SRoomPlaying",
}

var RoomState_value = map[string]int32{
	"SRoomIdle":    0,
	"SRoomWaiting": 1,
	"SRoomPlaying": 2,
}

func (x RoomState) Enum() *RoomState {
	p := new(RoomState)
	*p = x
	return p
}

func (x RoomState) String() string {
	return proto.EnumName(RoomState_name, int32(x))
}

func (x *RoomState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RoomState_value, data, "RoomState")
	if err != nil {
		return err
	}
	*x = RoomState(value)
	return nil
}

func (RoomState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{4}
}

type SSMsgBag struct {
	// 消息类型，SSMsgType
	MsgType *int32 `protobuf:"varint,1,req,name=msgType" json:"msgType,omitempty"`
	// 流水号，请求者初始化，回复者回填
	SeqNO *uint32 `protobuf:"varint,2,req,name=seqNO" json:"seqNO,omitempty"`
	// 请求消息码 SSMsgReqCode
	RequestCode *int32 `protobuf:"varint,3,req,name=requestCode" json:"requestCode,omitempty"`
	// 请求状态
	Status *int32 `protobuf:"varint,4,req,name=status" json:"status,omitempty"`
	// 额外参数
	Params []byte `protobuf:"bytes,5,opt,name=params" json:"params,omitempty"`
	// 源服务器URL，回复消息时使用
	SourceURL            *string  `protobuf:"bytes,6,opt,name=sourceURL" json:"sourceURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgBag) Reset()         { *m = SSMsgBag{} }
func (m *SSMsgBag) String() string { return proto.CompactTextString(m) }
func (*SSMsgBag) ProtoMessage()    {}
func (*SSMsgBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{0}
}
func (m *SSMsgBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgBag.Unmarshal(m, b)
}
func (m *SSMsgBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgBag.Marshal(b, m, deterministic)
}
func (m *SSMsgBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgBag.Merge(m, src)
}
func (m *SSMsgBag) XXX_Size() int {
	return xxx_messageInfo_SSMsgBag.Size(m)
}
func (m *SSMsgBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgBag.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgBag proto.InternalMessageInfo

func (m *SSMsgBag) GetMsgType() int32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *SSMsgBag) GetSeqNO() uint32 {
	if m != nil && m.SeqNO != nil {
		return *m.SeqNO
	}
	return 0
}

func (m *SSMsgBag) GetRequestCode() int32 {
	if m != nil && m.RequestCode != nil {
		return *m.RequestCode
	}
	return 0
}

func (m *SSMsgBag) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *SSMsgBag) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SSMsgBag) GetSourceURL() string {
	if m != nil && m.SourceURL != nil {
		return *m.SourceURL
	}
	return ""
}

// 房间管理服务器发送给游戏服务器
// 请求创建房间
type SSMsgCreateRoom struct {
	RoomID               *string  `protobuf:"bytes,1,req,name=roomID" json:"roomID,omitempty"`
	RoomConfigID         *string  `protobuf:"bytes,2,req,name=roomConfigID" json:"roomConfigID,omitempty"`
	RoomType             *int32   `protobuf:"varint,3,req,name=roomType" json:"roomType,omitempty"`
	UserID               *string  `protobuf:"bytes,4,req,name=userID" json:"userID,omitempty"`
	RoomNumber           *string  `protobuf:"bytes,5,req,name=roomNumber" json:"roomNumber,omitempty"`
	ClubID               *string  `protobuf:"bytes,6,opt,name=clubID" json:"clubID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgCreateRoom) Reset()         { *m = SSMsgCreateRoom{} }
func (m *SSMsgCreateRoom) String() string { return proto.CompactTextString(m) }
func (*SSMsgCreateRoom) ProtoMessage()    {}
func (*SSMsgCreateRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{1}
}
func (m *SSMsgCreateRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgCreateRoom.Unmarshal(m, b)
}
func (m *SSMsgCreateRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgCreateRoom.Marshal(b, m, deterministic)
}
func (m *SSMsgCreateRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgCreateRoom.Merge(m, src)
}
func (m *SSMsgCreateRoom) XXX_Size() int {
	return xxx_messageInfo_SSMsgCreateRoom.Size(m)
}
func (m *SSMsgCreateRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgCreateRoom.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgCreateRoom proto.InternalMessageInfo

func (m *SSMsgCreateRoom) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *SSMsgCreateRoom) GetRoomConfigID() string {
	if m != nil && m.RoomConfigID != nil {
		return *m.RoomConfigID
	}
	return ""
}

func (m *SSMsgCreateRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *SSMsgCreateRoom) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *SSMsgCreateRoom) GetRoomNumber() string {
	if m != nil && m.RoomNumber != nil {
		return *m.RoomNumber
	}
	return ""
}

func (m *SSMsgCreateRoom) GetClubID() string {
	if m != nil && m.ClubID != nil {
		return *m.ClubID
	}
	return ""
}

// 房间管理服务器发送给游戏服务器
// 请求删除房间
type SSMsgDeleteRoom struct {
	RoomID               *string  `protobuf:"bytes,1,req,name=roomID" json:"roomID,omitempty"`
	Why                  *int32   `protobuf:"varint,2,opt,name=why" json:"why,omitempty"`
	OnlyEmpty            *bool    `protobuf:"varint,3,opt,name=onlyEmpty" json:"onlyEmpty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgDeleteRoom) Reset()         { *m = SSMsgDeleteRoom{} }
func (m *SSMsgDeleteRoom) String() string { return proto.CompactTextString(m) }
func (*SSMsgDeleteRoom) ProtoMessage()    {}
func (*SSMsgDeleteRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{2}
}
func (m *SSMsgDeleteRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgDeleteRoom.Unmarshal(m, b)
}
func (m *SSMsgDeleteRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgDeleteRoom.Marshal(b, m, deterministic)
}
func (m *SSMsgDeleteRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgDeleteRoom.Merge(m, src)
}
func (m *SSMsgDeleteRoom) XXX_Size() int {
	return xxx_messageInfo_SSMsgDeleteRoom.Size(m)
}
func (m *SSMsgDeleteRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgDeleteRoom.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgDeleteRoom proto.InternalMessageInfo

func (m *SSMsgDeleteRoom) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *SSMsgDeleteRoom) GetWhy() int32 {
	if m != nil && m.Why != nil {
		return *m.Why
	}
	return 0
}

func (m *SSMsgDeleteRoom) GetOnlyEmpty() bool {
	if m != nil && m.OnlyEmpty != nil {
		return *m.OnlyEmpty
	}
	return false
}

// 用于用户ID列表，主要用于把用户列表保存到redis上
type SSMsgUserIDList struct {
	UserIDs              []string `protobuf:"bytes,1,rep,name=userIDs" json:"userIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgUserIDList) Reset()         { *m = SSMsgUserIDList{} }
func (m *SSMsgUserIDList) String() string { return proto.CompactTextString(m) }
func (*SSMsgUserIDList) ProtoMessage()    {}
func (*SSMsgUserIDList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{3}
}
func (m *SSMsgUserIDList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgUserIDList.Unmarshal(m, b)
}
func (m *SSMsgUserIDList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgUserIDList.Marshal(b, m, deterministic)
}
func (m *SSMsgUserIDList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgUserIDList.Merge(m, src)
}
func (m *SSMsgUserIDList) XXX_Size() int {
	return xxx_messageInfo_SSMsgUserIDList.Size(m)
}
func (m *SSMsgUserIDList) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgUserIDList.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgUserIDList proto.InternalMessageInfo

func (m *SSMsgUserIDList) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type SSMsgGameOverPlayerStat struct {
	UserID               *string  `protobuf:"bytes,1,req,name=userID" json:"userID,omitempty"`
	Score                *int32   `protobuf:"varint,2,req,name=score" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgGameOverPlayerStat) Reset()         { *m = SSMsgGameOverPlayerStat{} }
func (m *SSMsgGameOverPlayerStat) String() string { return proto.CompactTextString(m) }
func (*SSMsgGameOverPlayerStat) ProtoMessage()    {}
func (*SSMsgGameOverPlayerStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{4}
}
func (m *SSMsgGameOverPlayerStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgGameOverPlayerStat.Unmarshal(m, b)
}
func (m *SSMsgGameOverPlayerStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgGameOverPlayerStat.Marshal(b, m, deterministic)
}
func (m *SSMsgGameOverPlayerStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgGameOverPlayerStat.Merge(m, src)
}
func (m *SSMsgGameOverPlayerStat) XXX_Size() int {
	return xxx_messageInfo_SSMsgGameOverPlayerStat.Size(m)
}
func (m *SSMsgGameOverPlayerStat) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgGameOverPlayerStat.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgGameOverPlayerStat proto.InternalMessageInfo

func (m *SSMsgGameOverPlayerStat) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *SSMsgGameOverPlayerStat) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// 游戏服务器请求房间管理服务器删除一个房间
type SSMsgGameServer2RoomMgrServerDisbandRoom struct {
	RoomID               *string                    `protobuf:"bytes,1,req,name=roomID" json:"roomID,omitempty"`
	HandStart            *int32                     `protobuf:"varint,2,req,name=handStart" json:"handStart,omitempty"`
	HandFinished         *int32                     `protobuf:"varint,3,opt,name=handFinished" json:"handFinished,omitempty"`
	PlayerUserIDs        []string                   `protobuf:"bytes,4,rep,name=playerUserIDs" json:"playerUserIDs,omitempty"`
	PlayerStats          []*SSMsgGameOverPlayerStat `protobuf:"bytes,5,rep,name=playerStats" json:"playerStats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) Reset() {
	*m = SSMsgGameServer2RoomMgrServerDisbandRoom{}
}
func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) String() string { return proto.CompactTextString(m) }
func (*SSMsgGameServer2RoomMgrServerDisbandRoom) ProtoMessage()    {}
func (*SSMsgGameServer2RoomMgrServerDisbandRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{5}
}
func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgGameServer2RoomMgrServerDisbandRoom.Unmarshal(m, b)
}
func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgGameServer2RoomMgrServerDisbandRoom.Marshal(b, m, deterministic)
}
func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgGameServer2RoomMgrServerDisbandRoom.Merge(m, src)
}
func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) XXX_Size() int {
	return xxx_messageInfo_SSMsgGameServer2RoomMgrServerDisbandRoom.Size(m)
}
func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgGameServer2RoomMgrServerDisbandRoom.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgGameServer2RoomMgrServerDisbandRoom proto.InternalMessageInfo

func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) GetHandStart() int32 {
	if m != nil && m.HandStart != nil {
		return *m.HandStart
	}
	return 0
}

func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) GetHandFinished() int32 {
	if m != nil && m.HandFinished != nil {
		return *m.HandFinished
	}
	return 0
}

func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) GetPlayerUserIDs() []string {
	if m != nil {
		return m.PlayerUserIDs
	}
	return nil
}

func (m *SSMsgGameServer2RoomMgrServerDisbandRoom) GetPlayerStats() []*SSMsgGameOverPlayerStat {
	if m != nil {
		return m.PlayerStats
	}
	return nil
}

// 游戏服务器推送通知给房间管理服务器
// 通知房间状态变化
type SSMsgRoomStateNotify struct {
	State                *int32   `protobuf:"varint,1,req,name=state" json:"state,omitempty"`
	RoomID               *string  `protobuf:"bytes,2,req,name=roomID" json:"roomID,omitempty"`
	UserIDs              []string `protobuf:"bytes,3,rep,name=userIDs" json:"userIDs,omitempty"`
	HandStartted         *int32   `protobuf:"varint,4,req,name=handStartted" json:"handStartted,omitempty"`
	LastActiveTime       *uint32  `protobuf:"varint,5,req,name=lastActiveTime" json:"lastActiveTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgRoomStateNotify) Reset()         { *m = SSMsgRoomStateNotify{} }
func (m *SSMsgRoomStateNotify) String() string { return proto.CompactTextString(m) }
func (*SSMsgRoomStateNotify) ProtoMessage()    {}
func (*SSMsgRoomStateNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{6}
}
func (m *SSMsgRoomStateNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgRoomStateNotify.Unmarshal(m, b)
}
func (m *SSMsgRoomStateNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgRoomStateNotify.Marshal(b, m, deterministic)
}
func (m *SSMsgRoomStateNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgRoomStateNotify.Merge(m, src)
}
func (m *SSMsgRoomStateNotify) XXX_Size() int {
	return xxx_messageInfo_SSMsgRoomStateNotify.Size(m)
}
func (m *SSMsgRoomStateNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgRoomStateNotify.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgRoomStateNotify proto.InternalMessageInfo

func (m *SSMsgRoomStateNotify) GetState() int32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SSMsgRoomStateNotify) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *SSMsgRoomStateNotify) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

func (m *SSMsgRoomStateNotify) GetHandStartted() int32 {
	if m != nil && m.HandStartted != nil {
		return *m.HandStartted
	}
	return 0
}

func (m *SSMsgRoomStateNotify) GetLastActiveTime() uint32 {
	if m != nil && m.LastActiveTime != nil {
		return *m.LastActiveTime
	}
	return 0
}

// 游戏服务器请求房间服务器扣钱或者返还钱
// 扣钱或者返还钱参见SSMsgReqCode
type SSMsgUpdateBalance struct {
	UserID               *string  `protobuf:"bytes,1,req,name=userID" json:"userID,omitempty"`
	RoomID               *string  `protobuf:"bytes,2,req,name=roomID" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgUpdateBalance) Reset()         { *m = SSMsgUpdateBalance{} }
func (m *SSMsgUpdateBalance) String() string { return proto.CompactTextString(m) }
func (*SSMsgUpdateBalance) ProtoMessage()    {}
func (*SSMsgUpdateBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{7}
}
func (m *SSMsgUpdateBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgUpdateBalance.Unmarshal(m, b)
}
func (m *SSMsgUpdateBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgUpdateBalance.Marshal(b, m, deterministic)
}
func (m *SSMsgUpdateBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgUpdateBalance.Merge(m, src)
}
func (m *SSMsgUpdateBalance) XXX_Size() int {
	return xxx_messageInfo_SSMsgUpdateBalance.Size(m)
}
func (m *SSMsgUpdateBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgUpdateBalance.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgUpdateBalance proto.InternalMessageInfo

func (m *SSMsgUpdateBalance) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *SSMsgUpdateBalance) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

// 道具打赏
type SSMsgDonate struct {
	PropsType            *int32   `protobuf:"varint,1,req,name=propsType" json:"propsType,omitempty"`
	From                 *string  `protobuf:"bytes,2,req,name=from" json:"from,omitempty"`
	To                   *string  `protobuf:"bytes,3,req,name=to" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgDonate) Reset()         { *m = SSMsgDonate{} }
func (m *SSMsgDonate) String() string { return proto.CompactTextString(m) }
func (*SSMsgDonate) ProtoMessage()    {}
func (*SSMsgDonate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{8}
}
func (m *SSMsgDonate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgDonate.Unmarshal(m, b)
}
func (m *SSMsgDonate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgDonate.Marshal(b, m, deterministic)
}
func (m *SSMsgDonate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgDonate.Merge(m, src)
}
func (m *SSMsgDonate) XXX_Size() int {
	return xxx_messageInfo_SSMsgDonate.Size(m)
}
func (m *SSMsgDonate) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgDonate.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgDonate proto.InternalMessageInfo

func (m *SSMsgDonate) GetPropsType() int32 {
	if m != nil && m.PropsType != nil {
		return *m.PropsType
	}
	return 0
}

func (m *SSMsgDonate) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *SSMsgDonate) GetTo() string {
	if m != nil && m.To != nil {
		return *m.To
	}
	return ""
}

// 房间服务器返回用户钻石与魅力值给游戏服务器
type SSMsgDonateRsp struct {
	Diamond              *int32   `protobuf:"varint,1,req,name=diamond" json:"diamond,omitempty"`
	Charm                *int32   `protobuf:"varint,2,req,name=charm" json:"charm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgDonateRsp) Reset()         { *m = SSMsgDonateRsp{} }
func (m *SSMsgDonateRsp) String() string { return proto.CompactTextString(m) }
func (*SSMsgDonateRsp) ProtoMessage()    {}
func (*SSMsgDonateRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{9}
}
func (m *SSMsgDonateRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgDonateRsp.Unmarshal(m, b)
}
func (m *SSMsgDonateRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgDonateRsp.Marshal(b, m, deterministic)
}
func (m *SSMsgDonateRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgDonateRsp.Merge(m, src)
}
func (m *SSMsgDonateRsp) XXX_Size() int {
	return xxx_messageInfo_SSMsgDonateRsp.Size(m)
}
func (m *SSMsgDonateRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgDonateRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgDonateRsp proto.InternalMessageInfo

func (m *SSMsgDonateRsp) GetDiamond() int32 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *SSMsgDonateRsp) GetCharm() int32 {
	if m != nil && m.Charm != nil {
		return *m.Charm
	}
	return 0
}

// 更新用户的位置信息到房间里
type SSMsgUpdateLocation struct {
	UserID               *string  `protobuf:"bytes,1,req,name=userID" json:"userID,omitempty"`
	Location             *string  `protobuf:"bytes,2,req,name=location" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgUpdateLocation) Reset()         { *m = SSMsgUpdateLocation{} }
func (m *SSMsgUpdateLocation) String() string { return proto.CompactTextString(m) }
func (*SSMsgUpdateLocation) ProtoMessage()    {}
func (*SSMsgUpdateLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{10}
}
func (m *SSMsgUpdateLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgUpdateLocation.Unmarshal(m, b)
}
func (m *SSMsgUpdateLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgUpdateLocation.Marshal(b, m, deterministic)
}
func (m *SSMsgUpdateLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgUpdateLocation.Merge(m, src)
}
func (m *SSMsgUpdateLocation) XXX_Size() int {
	return xxx_messageInfo_SSMsgUpdateLocation.Size(m)
}
func (m *SSMsgUpdateLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgUpdateLocation.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgUpdateLocation proto.InternalMessageInfo

func (m *SSMsgUpdateLocation) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *SSMsgUpdateLocation) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

// 游戏服务器推送通知给房间管理服务器
// 通知收牌开始
type SSMsgHandBeginNotify struct {
	RoomID               *string  `protobuf:"bytes,2,req,name=roomID" json:"roomID,omitempty"`
	HandStartted         *int32   `protobuf:"varint,4,req,name=handStartted" json:"handStartted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSMsgHandBeginNotify) Reset()         { *m = SSMsgHandBeginNotify{} }
func (m *SSMsgHandBeginNotify) String() string { return proto.CompactTextString(m) }
func (*SSMsgHandBeginNotify) ProtoMessage()    {}
func (*SSMsgHandBeginNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25bd73e18305d7e, []int{11}
}
func (m *SSMsgHandBeginNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSMsgHandBeginNotify.Unmarshal(m, b)
}
func (m *SSMsgHandBeginNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSMsgHandBeginNotify.Marshal(b, m, deterministic)
}
func (m *SSMsgHandBeginNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSMsgHandBeginNotify.Merge(m, src)
}
func (m *SSMsgHandBeginNotify) XXX_Size() int {
	return xxx_messageInfo_SSMsgHandBeginNotify.Size(m)
}
func (m *SSMsgHandBeginNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_SSMsgHandBeginNotify.DiscardUnknown(m)
}

var xxx_messageInfo_SSMsgHandBeginNotify proto.InternalMessageInfo

func (m *SSMsgHandBeginNotify) GetRoomID() string {
	if m != nil && m.RoomID != nil {
		return *m.RoomID
	}
	return ""
}

func (m *SSMsgHandBeginNotify) GetHandStartted() int32 {
	if m != nil && m.HandStartted != nil {
		return *m.HandStartted
	}
	return 0
}

func init() {
	proto.RegisterType((*SSMsgBag)(nil), "gconst.SSMsgBag")
	proto.RegisterType((*SSMsgCreateRoom)(nil), "gconst.SSMsgCreateRoom")
	proto.RegisterType((*SSMsgDeleteRoom)(nil), "gconst.SSMsgDeleteRoom")
	proto.RegisterType((*SSMsgUserIDList)(nil), "gconst.SSMsgUserIDList")
	proto.RegisterType((*SSMsgGameOverPlayerStat)(nil), "gconst.SSMsgGameOverPlayerStat")
	proto.RegisterType((*SSMsgGameServer2RoomMgrServerDisbandRoom)(nil), "gconst.SSMsgGameServer2RoomMgrServerDisbandRoom")
	proto.RegisterType((*SSMsgRoomStateNotify)(nil), "gconst.SSMsgRoomStateNotify")
	proto.RegisterType((*SSMsgUpdateBalance)(nil), "gconst.SSMsgUpdateBalance")
	proto.RegisterType((*SSMsgDonate)(nil), "gconst.SSMsgDonate")
	proto.RegisterType((*SSMsgDonateRsp)(nil), "gconst.SSMsgDonateRsp")
	proto.RegisterType((*SSMsgUpdateLocation)(nil), "gconst.SSMsgUpdateLocation")
	proto.RegisterType((*SSMsgHandBeginNotify)(nil), "gconst.SSMsgHandBeginNotify")
	proto.RegisterEnum("gconst.SSMsgType", SSMsgType_name, SSMsgType_value)
	proto.RegisterEnum("gconst.SSMsgReqCode", SSMsgReqCode_name, SSMsgReqCode_value)
	proto.RegisterEnum("gconst.SSMsgError", SSMsgError_name, SSMsgError_value)
	proto.RegisterEnum("gconst.RoomType", RoomType_name, RoomType_value)
	proto.RegisterEnum("gconst.RoomState", RoomState_name, RoomState_value)
}

func init() { proto.RegisterFile("gconst.proto", fileDescriptor_a25bd73e18305d7e) }

var fileDescriptor_a25bd73e18305d7e = []byte{
	// 1035 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4f, 0x6f, 0xdb, 0xc6,
	0x13, 0x0d, 0xa9, 0xbf, 0x1c, 0xfd, 0xf1, 0x7a, 0xed, 0xfc, 0xa2, 0xdf, 0x21, 0xa8, 0x2a, 0xf4,
	0x20, 0xf8, 0x10, 0x14, 0x6e, 0x8a, 0x14, 0x45, 0x51, 0x40, 0x31, 0x69, 0x5b, 0xaa, 0x25, 0x1b,
	0x92, 0x8d, 0x26, 0xba, 0xad, 0xc9, 0x15, 0x45, 0x44, 0xda, 0xa5, 0x77, 0x97, 0x4e, 0x74, 0xe8,
	0x97, 0xe9, 0xa5, 0xd7, 0xa2, 0x9f, 0xb0, 0x98, 0x25, 0x23, 0xab, 0x4e, 0x7c, 0xe3, 0x9b, 0xdd,
	0x99, 0xf7, 0xe6, 0xcd, 0x2c, 0xa1, 0x19, 0x87, 0x52, 0x68, 0xf3, 0x2a, 0x55, 0xd2, 0x48, 0x5a,
	0xcd, 0x51, 0x4f, 0x41, 0x7d, 0x36, 0x1b, 0xeb, 0xf8, 0x2d, 0x8b, 0xe9, 0x1e, 0xd4, 0xd6, 0x3a,
	0xbe, 0xde, 0xa4, 0xbc, 0xe3, 0x74, 0xdd, 0x7e, 0x85, 0xb6, 0xa0, 0xa2, 0xf9, 0xdd, 0xe4, 0xb2,
	0xe3, 0x76, 0xdd, 0x7e, 0x8b, 0x1e, 0x40, 0x43, 0xf1, 0xbb, 0x8c, 0x6b, 0x73, 0x22, 0x23, 0xde,
	0x29, 0xd9, 0x3b, 0x6d, 0xa8, 0x6a, 0xc3, 0x4c, 0xa6, 0x3b, 0xe5, 0xcf, 0x38, 0x65, 0x8a, 0xad,
	0x75, 0xa7, 0xd2, 0x75, 0xfa, 0x4d, 0xba, 0x0f, 0x9e, 0x96, 0x99, 0x0a, 0xf9, 0xcd, 0xf4, 0xa2,
	0x53, 0xed, 0x3a, 0x7d, 0xaf, 0xf7, 0x07, 0xec, 0x59, 0xce, 0x13, 0xc5, 0x99, 0xe1, 0x53, 0x29,
	0xd7, 0x98, 0xa5, 0xa4, 0x5c, 0x0f, 0x7d, 0xcb, 0xec, 0xd1, 0x43, 0x68, 0x22, 0x3e, 0x91, 0x62,
	0x91, 0xc4, 0x43, 0xdf, 0x0a, 0xf0, 0x28, 0x81, 0x3a, 0x46, 0xad, 0xc2, 0x2d, 0x7b, 0xa6, 0xb9,
	0x1a, 0xfa, 0x96, 0xdd, 0xa3, 0x14, 0x00, 0x6f, 0x4c, 0xb2, 0xf5, 0x2d, 0x57, 0x9d, 0x8a, 0x8d,
	0xb5, 0xa1, 0x1a, 0xae, 0xb2, 0xdb, 0xa1, 0x5f, 0xd0, 0x0f, 0x0a, 0x7a, 0x9f, 0xaf, 0xf8, 0x13,
	0xf4, 0x0d, 0x28, 0x7d, 0x5c, 0x6e, 0x3a, 0x6e, 0xd7, 0xe9, 0x57, 0xb0, 0x03, 0x29, 0x56, 0x9b,
	0x60, 0x9d, 0x9a, 0x4d, 0xa7, 0xd4, 0x75, 0xfa, 0xf5, 0x5e, 0xaf, 0x28, 0x71, 0x63, 0xb9, 0x2f,
	0x12, 0x6d, 0xd0, 0xbc, 0x5c, 0x89, 0xee, 0x38, 0xdd, 0x52, 0xdf, 0xeb, 0xfd, 0x04, 0x2f, 0xec,
	0x9d, 0x33, 0xb6, 0xe6, 0x97, 0xf7, 0x5c, 0x5d, 0xad, 0xd8, 0x86, 0xab, 0x99, 0x61, 0x66, 0x47,
	0x75, 0x4e, 0x87, 0x3e, 0x87, 0x52, 0x71, 0xdb, 0x66, 0xa5, 0xf7, 0xb7, 0x03, 0xfd, 0x6d, 0xea,
	0x8c, 0xab, 0x7b, 0xae, 0x8e, 0x51, 0xe6, 0x38, 0x56, 0x39, 0xf2, 0x13, 0x7d, 0xcb, 0x44, 0xf4,
	0x55, 0xe9, 0xfb, 0xe0, 0x2d, 0x99, 0x88, 0x66, 0x86, 0x29, 0x93, 0xd7, 0x43, 0x33, 0x31, 0x74,
	0x9a, 0x88, 0x44, 0x2f, 0x79, 0x64, 0x7b, 0xa8, 0xd0, 0xe7, 0xd0, 0x4a, 0xad, 0xa4, 0x9b, 0x42,
	0x76, 0x19, 0x65, 0xd3, 0xd7, 0xd0, 0x48, 0xb7, 0x4a, 0x71, 0x88, 0xa5, 0x7e, 0xe3, 0xf8, 0x9b,
	0x57, 0xc5, 0xf2, 0x3c, 0xd1, 0x51, 0xcf, 0xc0, 0xa1, 0x3d, 0x42, 0x49, 0x18, 0xe0, 0x13, 0x69,
	0x92, 0xc5, 0xc6, 0x76, 0x86, 0xb0, 0x58, 0xa8, 0x07, 0xb1, 0xf9, 0x40, 0x77, 0x4c, 0x2b, 0x59,
	0xf6, 0x42, 0xaa, 0x55, 0x6f, 0x78, 0x54, 0xec, 0xd4, 0xff, 0xa0, 0xbd, 0x62, 0xda, 0x0c, 0x42,
	0x93, 0xdc, 0xf3, 0xeb, 0x64, 0xcd, 0xed, 0x64, 0x5b, 0xbd, 0xd7, 0x40, 0xf3, 0x31, 0xa4, 0x11,
	0x33, 0xfc, 0x2d, 0x5b, 0x31, 0x11, 0xf2, 0x2f, 0xdc, 0x7d, 0x44, 0xda, 0xfb, 0x19, 0x1a, 0xf9,
	0xfc, 0xa5, 0x60, 0x86, 0xa3, 0x61, 0xa9, 0x92, 0xa9, 0xde, 0xd9, 0xfb, 0x26, 0x94, 0x17, 0x4a,
	0xae, 0x0b, 0x91, 0x00, 0xae, 0x91, 0x76, 0xdf, 0xbc, 0xde, 0xf7, 0xd0, 0xde, 0xc9, 0x9d, 0xea,
	0x14, 0x5b, 0x88, 0x12, 0xb6, 0x96, 0x22, 0x7a, 0x78, 0x34, 0xe1, 0x92, 0xa9, 0x75, 0x31, 0xcc,
	0x37, 0x70, 0xb0, 0xa3, 0xf1, 0x42, 0x86, 0xcc, 0x24, 0x52, 0x7c, 0x21, 0x92, 0x40, 0x7d, 0x55,
	0x9c, 0x15, 0x32, 0x7f, 0x29, 0x2c, 0x3d, 0x67, 0x22, 0x7a, 0xcb, 0xe3, 0x44, 0x14, 0x96, 0x3e,
	0xf6, 0xf0, 0xab, 0x96, 0x1d, 0x1d, 0x83, 0x67, 0xb3, 0xb1, 0x2b, 0xda, 0x80, 0xda, 0x34, 0x7f,
	0xb8, 0xc4, 0xa1, 0x4d, 0xa8, 0x4f, 0xb9, 0x4e, 0xa5, 0xd0, 0x9c, 0xb8, 0x14, 0xa0, 0x9a, 0xd7,
	0x25, 0xa5, 0xa3, 0x7f, 0x1c, 0x68, 0xe6, 0x53, 0xe4, 0x77, 0xf8, 0xc2, 0x69, 0x1b, 0xe0, 0xe1,
	0x8d, 0x12, 0x07, 0xf1, 0xc3, 0xa3, 0x21, 0x2e, 0x3d, 0x80, 0xbd, 0x47, 0x03, 0x27, 0x25, 0x7a,
	0x08, 0x64, 0x30, 0x08, 0x3e, 0x25, 0x06, 0x8f, 0x8a, 0x68, 0x99, 0xee, 0x41, 0x63, 0x30, 0x08,
	0x84, 0xe1, 0xca, 0xe6, 0x56, 0x90, 0x38, 0x37, 0x91, 0x54, 0x29, 0x85, 0xf6, 0x7f, 0xed, 0x21,
	0x35, 0xba, 0x0f, 0xad, 0x3c, 0x76, 0xa5, 0x64, 0x7a, 0xb2, 0x88, 0x49, 0x1d, 0xe9, 0x1e, 0x99,
	0x41, 0xbc, 0xa3, 0xbf, 0x5c, 0x00, 0x2b, 0x3a, 0x50, 0x4a, 0x2a, 0x94, 0x18, 0x28, 0x35, 0xcb,
	0xc2, 0x90, 0x6b, 0x4d, 0x9e, 0xd1, 0x16, 0x78, 0x81, 0x52, 0x3e, 0x0f, 0x65, 0xc4, 0x89, 0x53,
	0xc0, 0x40, 0x58, 0xe8, 0x52, 0x02, 0xcd, 0x40, 0x59, 0x45, 0xc1, 0xa7, 0x44, 0x1b, 0x52, 0x42,
	0x8e, 0x40, 0xa9, 0x89, 0x9c, 0x6e, 0xff, 0x3e, 0xa4, 0x4c, 0x3b, 0x70, 0x18, 0x28, 0x75, 0x23,
	0x74, 0x96, 0xa6, 0x52, 0xd9, 0xc6, 0xd0, 0x57, 0x52, 0xa1, 0x2f, 0xe0, 0x60, 0x5b, 0x7e, 0x27,
	0xa5, 0x5a, 0xd4, 0x29, 0x2c, 0xc8, 0x8b, 0xd7, 0xe8, 0x77, 0xd0, 0x0d, 0x94, 0xba, 0x66, 0x1f,
	0xb8, 0x5c, 0x2c, 0xfc, 0x7c, 0x6b, 0x4e, 0x59, 0xb2, 0xe2, 0x11, 0x5e, 0x12, 0x32, 0x8b, 0x97,
	0xc4, 0xa3, 0x2f, 0xe1, 0xff, 0x4f, 0xdc, 0x1a, 0x5e, 0x12, 0xa0, 0xdf, 0xc2, 0xcb, 0x27, 0x8e,
	0xa7, 0x3c, 0xe5, 0xcc, 0x90, 0x06, 0x8e, 0xa0, 0x20, 0x1f, 0xea, 0x89, 0xb4, 0x3f, 0x2e, 0xd2,
	0x3c, 0xfa, 0xd3, 0x85, 0xfa, 0x67, 0xe9, 0xb8, 0x05, 0x3e, 0x5b, 0x70, 0x11, 0x8f, 0x47, 0xf9,
	0x4e, 0x9c, 0x65, 0x4c, 0xf8, 0x4c, 0xc4, 0xc4, 0x45, 0x93, 0x7c, 0x29, 0xe2, 0x6b, 0x96, 0x8c,
	0x47, 0xa4, 0x84, 0x96, 0xbe, 0x67, 0xe2, 0x64, 0x99, 0x5f, 0x2e, 0x23, 0x9e, 0x2d, 0x99, 0xc4,
	0x84, 0xf1, 0x88, 0xe0, 0x6b, 0xa9, 0x4f, 0x12, 0x11, 0x0f, 0x10, 0x55, 0xf1, 0xf4, 0x5d, 0x22,
	0x46, 0x09, 0x13, 0xf1, 0xd9, 0x39, 0xa9, 0x3d, 0x10, 0x9d, 0xcd, 0x49, 0x1d, 0xd7, 0x20, 0x47,
	0x6f, 0x3e, 0xfe, 0x78, 0xfc, 0x03, 0xf1, 0xec, 0xfe, 0x25, 0xd9, 0x24, 0xc9, 0x08, 0xd0, 0x1a,
	0x94, 0x7c, 0x7f, 0x4e, 0x1a, 0x98, 0xf3, 0x2e, 0xe3, 0x17, 0x49, 0x36, 0x1e, 0x91, 0x26, 0xca,
	0xb9, 0x60, 0x62, 0xbe, 0x94, 0x08, 0x5b, 0x16, 0x22, 0xb6, 0xb0, 0x8d, 0x70, 0xbe, 0x64, 0x22,
	0x7e, 0xcf, 0xc7, 0x23, 0xb2, 0x87, 0xf5, 0xae, 0x59, 0x88, 0x52, 0x08, 0xae, 0x3d, 0x7e, 0x5f,
	0xc9, 0x0f, 0x64, 0x1f, 0x8b, 0x5f, 0x9d, 0xcf, 0x09, 0xa5, 0x75, 0x28, 0x9f, 0xfe, 0x36, 0x1e,
	0x91, 0x03, 0xfb, 0x75, 0x33, 0x1e, 0x91, 0x43, 0xfc, 0x9a, 0x8f, 0xc6, 0x23, 0xf2, 0xfc, 0xe8,
	0x57, 0xf0, 0xb6, 0x2b, 0x8d, 0xb5, 0x67, 0xd6, 0xc5, 0x68, 0xc5, 0xc9, 0x33, 0xdc, 0x16, 0x0b,
	0x7f, 0x67, 0x89, 0x49, 0x44, 0x4c, 0x9c, 0x6d, 0x04, 0xff, 0x84, 0x18, 0x71, 0xff, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0x4c, 0x76, 0x03, 0x66, 0x07, 0x00, 0x00,
}
