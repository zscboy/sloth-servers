// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_pokerface_ddz.proto

/*
Package pddz is a generated protocol buffer package.

It is generated from these files:
	game_pokerface_ddz.proto

It has these top-level messages:
*/
package pddz

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 牌组类型
type CardHandType int32

const (
	CardHandType_None             CardHandType = 0
	CardHandType_Flush            CardHandType = 1
	CardHandType_Bomb             CardHandType = 2
	CardHandType_Single           CardHandType = 3
	CardHandType_Pair             CardHandType = 4
	CardHandType_Pair3X           CardHandType = 5
	CardHandType_Triplet          CardHandType = 6
	CardHandType_TripletSingle    CardHandType = 7
	CardHandType_TripletPair      CardHandType = 8
	CardHandType_Triplet2X        CardHandType = 9
	CardHandType_Triplet2X2Pair   CardHandType = 10
	CardHandType_Triplet2X2Single CardHandType = 11
	CardHandType_FourX2Single     CardHandType = 12
	CardHandType_FourX2Pair       CardHandType = 13
	CardHandType_Roket            CardHandType = 14
)

var CardHandType_name = map[int32]string{
	0:  "None",
	1:  "Flush",
	2:  "Bomb",
	3:  "Single",
	4:  "Pair",
	5:  "Pair3X",
	6:  "Triplet",
	7:  "TripletSingle",
	8:  "TripletPair",
	9:  "Triplet2X",
	10: "Triplet2X2Pair",
	11: "Triplet2X2Single",
	12: "FourX2Single",
	13: "FourX2Pair",
	14: "Roket",
}
var CardHandType_value = map[string]int32{
	"None":             0,
	"Flush":            1,
	"Bomb":             2,
	"Single":           3,
	"Pair":             4,
	"Pair3X":           5,
	"Triplet":          6,
	"TripletSingle":    7,
	"TripletPair":      8,
	"Triplet2X":        9,
	"Triplet2X2Pair":   10,
	"Triplet2X2Single": 11,
	"FourX2Single":     12,
	"FourX2Pair":       13,
	"Roket":            14,
}

func (x CardHandType) Enum() *CardHandType {
	p := new(CardHandType)
	*p = x
	return p
}
func (x CardHandType) String() string {
	return proto.EnumName(CardHandType_name, int32(x))
}
func (x *CardHandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CardHandType_value, data, "CardHandType")
	if err != nil {
		return err
	}
	*x = CardHandType(value)
	return nil
}
func (CardHandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 一手牌局结束
// 可能的结果是：流局、有人自摸胡牌、有人放铳其他人胡牌
type HandOverType int32

const (
	HandOverType_enumHandOverType_None          HandOverType = 0
	HandOverType_enumHandOverType_Win_SelfDrawn HandOverType = 1
	HandOverType_enumHandOverType_Win_Chuck     HandOverType = 2
	HandOverType_enumHandOverType_Chucker       HandOverType = 3
	HandOverType_enumHandOverType_Konger        HandOverType = 4
	HandOverType_enumHandOverType_Win_RobKong   HandOverType = 5
)

var HandOverType_name = map[int32]string{
	0: "enumHandOverType_None",
	1: "enumHandOverType_Win_SelfDrawn",
	2: "enumHandOverType_Win_Chuck",
	3: "enumHandOverType_Chucker",
	4: "enumHandOverType_Konger",
	5: "enumHandOverType_Win_RobKong",
}
var HandOverType_value = map[string]int32{
	"enumHandOverType_None":          0,
	"enumHandOverType_Win_SelfDrawn": 1,
	"enumHandOverType_Win_Chuck":     2,
	"enumHandOverType_Chucker":       3,
	"enumHandOverType_Konger":        4,
	"enumHandOverType_Win_RobKong":   5,
}

func (x HandOverType) Enum() *HandOverType {
	p := new(HandOverType)
	*p = x
	return p
}
func (x HandOverType) String() string {
	return proto.EnumName(HandOverType_name, int32(x))
}
func (x *HandOverType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HandOverType_value, data, "HandOverType")
	if err != nil {
		return err
	}
	*x = HandOverType(value)
	return nil
}
func (HandOverType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 动作类型
// 注意为了能够用一个int合并多个动作
// 因此所有动作的值均为二进制bit field独立
type ActionType int32

const (
	ActionType_enumActionType_None          ActionType = 0
	ActionType_enumActionType_SKIP          ActionType = 1
	ActionType_enumActionType_DISCARD       ActionType = 2
	ActionType_enumActionType_DRAW          ActionType = 4
	ActionType_enumActionType_Win_SelfDrawn ActionType = 8
	ActionType_enumActionType_Call          ActionType = 16
	ActionType_enumActionType_Rob           ActionType = 32
	ActionType_enumActionType_CallDouble    ActionType = 64
	ActionType_enumActionType_CallWithScore ActionType = 128
)

var ActionType_name = map[int32]string{
	0:   "enumActionType_None",
	1:   "enumActionType_SKIP",
	2:   "enumActionType_DISCARD",
	4:   "enumActionType_DRAW",
	8:   "enumActionType_Win_SelfDrawn",
	16:  "enumActionType_Call",
	32:  "enumActionType_Rob",
	64:  "enumActionType_CallDouble",
	128: "enumActionType_CallWithScore",
}
var ActionType_value = map[string]int32{
	"enumActionType_None":          0,
	"enumActionType_SKIP":          1,
	"enumActionType_DISCARD":       2,
	"enumActionType_DRAW":          4,
	"enumActionType_Win_SelfDrawn": 8,
	"enumActionType_Call":          16,
	"enumActionType_Rob":           32,
	"enumActionType_CallDouble":    64,
	"enumActionType_CallWithScore": 128,
}

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}
func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}
func (x *ActionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ActionType_value, data, "ActionType")
	if err != nil {
		return err
	}
	*x = ActionType(value)
	return nil
}
func (ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterEnum("pddz.CardHandType", CardHandType_name, CardHandType_value)
	proto.RegisterEnum("pddz.HandOverType", HandOverType_name, HandOverType_value)
	proto.RegisterEnum("pddz.ActionType", ActionType_name, ActionType_value)
}

func init() { proto.RegisterFile("game_pokerface_ddz.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0xc7, 0x6d, 0x2f, 0xdc, 0x0b, 0x87, 0x82, 0xc7, 0xa3, 0xde, 0x0f, 0x44, 0x82, 0x2e, 0x59,
	0xb8, 0xc0, 0x17, 0x10, 0xdb, 0x10, 0x09, 0x89, 0x92, 0x96, 0x04, 0x76, 0x4d, 0x3f, 0x06, 0x68,
	0x28, 0x33, 0xcd, 0xd8, 0x6a, 0x64, 0xe5, 0xd2, 0x67, 0xf3, 0x41, 0x7c, 0x0e, 0x33, 0xb5, 0x02,
	0x97, 0x76, 0xd7, 0xf9, 0xfd, 0xce, 0x7f, 0x3a, 0xff, 0x03, 0xf7, 0x1b, 0x6f, 0xcf, 0xdc, 0x44,
	0xec, 0x98, 0x5c, 0x7b, 0x01, 0x73, 0xc3, 0xf0, 0xf0, 0x2e, 0x91, 0x22, 0x15, 0x54, 0x4b, 0xc2,
	0xf0, 0x30, 0xfc, 0xa3, 0x81, 0x61, 0x7a, 0x32, 0xfc, 0xe4, 0xf1, 0x70, 0xf1, 0x23, 0x61, 0xd4,
	0x80, 0xda, 0x67, 0xc1, 0x19, 0x3e, 0xa1, 0x26, 0xd4, 0x27, 0x71, 0xf6, 0x75, 0x8b, 0x9a, 0x82,
	0x1f, 0xc5, 0xde, 0x47, 0x9d, 0x00, 0xae, 0x9d, 0x88, 0x6f, 0x62, 0x86, 0x57, 0x8a, 0xce, 0xbd,
	0x48, 0x62, 0x4d, 0x51, 0xf5, 0xf5, 0x7e, 0x85, 0x75, 0x6a, 0xc1, 0xcd, 0x42, 0x46, 0x49, 0xcc,
	0x52, 0xbc, 0xa6, 0x67, 0xd0, 0x2e, 0x0e, 0x45, 0xea, 0x86, 0x9e, 0x42, 0xab, 0x40, 0x79, 0xb8,
	0x41, 0x6d, 0x68, 0x16, 0x60, 0xb4, 0xc2, 0x26, 0x11, 0x74, 0x8e, 0xc7, 0x51, 0x3e, 0x02, 0xf4,
	0x02, 0xf0, 0xc4, 0x8a, 0x9b, 0x5a, 0x84, 0x60, 0x4c, 0x44, 0x26, 0x8f, 0xc4, 0xa0, 0x0e, 0xc0,
	0x3f, 0x92, 0xe7, 0xda, 0xaa, 0x82, 0x2d, 0x76, 0x2c, 0xc5, 0xce, 0xf0, 0xb7, 0x06, 0x86, 0x2a,
	0xf9, 0xe5, 0x1b, 0x93, 0x79, 0xd1, 0x07, 0x78, 0xc9, 0x78, 0xb6, 0x3f, 0x67, 0x6e, 0xd1, 0xfc,
	0x2d, 0xf4, 0x4b, 0x6a, 0x19, 0x71, 0xd7, 0x61, 0xf1, 0xda, 0x92, 0xde, 0x77, 0x8e, 0x1a, 0xf5,
	0xa1, 0x5b, 0x39, 0x63, 0x6e, 0xb3, 0x60, 0x87, 0x3a, 0xf5, 0xe0, 0xbe, 0xe4, 0x73, 0xc7, 0x24,
	0x5e, 0xd1, 0x2b, 0xb8, 0x2b, 0xd9, 0x99, 0xe0, 0x1b, 0xa6, 0xb6, 0x39, 0x80, 0x5e, 0xe5, 0xd5,
	0xb6, 0xf0, 0xd5, 0x0c, 0xd6, 0x87, 0xbf, 0x74, 0x80, 0x71, 0x90, 0x46, 0x82, 0xe7, 0x55, 0xee,
	0xe0, 0xb9, 0x0a, 0x9c, 0xc8, 0xff, 0x22, 0x65, 0xe1, 0xcc, 0xa6, 0x73, 0xd4, 0xa8, 0x0b, 0xb7,
	0x17, 0xc2, 0x9a, 0x3a, 0xe6, 0xd8, 0xb6, 0x50, 0xaf, 0x08, 0x59, 0xf6, 0x78, 0x79, 0x7a, 0xd7,
	0x99, 0x78, 0xbc, 0x94, 0x46, 0x45, 0xd4, 0xf4, 0xe2, 0x18, 0x91, 0x6e, 0x81, 0x2e, 0x84, 0x2d,
	0x7c, 0x1c, 0xd0, 0x6b, 0x78, 0xa8, 0x08, 0x58, 0x22, 0xf3, 0x63, 0x86, 0x1f, 0xe8, 0x4d, 0xe9,
	0x8f, 0x4a, 0x2f, 0xa3, 0x74, 0xeb, 0x04, 0x42, 0x32, 0xfc, 0xa9, 0xfd, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x09, 0x69, 0x2a, 0xe1, 0x02, 0x00, 0x00,
}
