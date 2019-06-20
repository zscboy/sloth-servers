package club

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"lobbyserver/lobby"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
)

// TestSomething 测试用例
func TestSomething(t *testing.T) {
	log.Println("TestSomething")
	log.Println("now time:", int64(time.Millisecond))
	// testCreateClub("10000002")
	// testLoadMyClubs("10000002")
	// testDeleteClub("10000002")
	// testLoadClubMembers("10000019")
	// testJoinClub("10000004", "56135")
	// testLoadClubEvent("10000002")
	// testJoinApproval("10000001", "10000004", "0aee5aa1-8765-11e9-9a1a-00163e0f7404", "yes", "11")
	// testClubQuit("10000003")
	// testLoadMyClubs("10000003")
	// testCreateClubRoom("10000002")
	// testLoadClubRooms("10000002")
	// testDeleteClubRoom("10000002")
	// testLoadMyApplyEvent("10000020")
	// renameClub("10000019", "4088fc70-8e50-11e9-8fea-107b445225b6", "哈哈哈")
	// kickOutMember("10000019", "83f6aa30-8e63-11e9-8fea-107b445225b6", "10000020")
	// changeClubMemberRole("10000019", "83f6aa30-8e63-11e9-8fea-107b445225b6", "10000020", int(ClubRoleType_CRoleTypeMgr))
	// loadClubManagers("10000019", "83f6aa30-8e63-11e9-8fea-107b445225b6")
	// allowMemberCreateRoom("10000019", "83f6aa30-8e63-11e9-8fea-107b445225b6", "10000020")
	testSetRoomOptions("10000019", "d521d820-9265-11e9-8fea-107b445225b6", "{\"abc\":\"aaa\"}")

}

func testCreateClub(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	var url = "http://localhost:3002/lobby/uuid/createClub?tk=" + tk + "&clname=mytest"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	createClubReply := &MsgCreateClubReply{}
	buf := msgClubReply.GetContent()

	err = proto.Unmarshal(buf, createClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	log.Println("createClubReply:", createClubReply)
}

func testLoadMyClubs(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	var url = "http://localhost:3002/lobby/uuid/loadMyClubs?tk=" + tk
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	reply := &MsgClubLoadMyClubsReply{}
	buf := msgClubReply.GetContent()

	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("err:", err)
	}

	log.Println("reply:", reply)
}

func testDeleteClub(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	var url = "http://localhost:3002/lobby/uuid/disbandClub?tk=" + tk + "&clubID=9949cd58-7e97-11e9-a192-107b445225b6"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	reply := &MsgClubLoadMyClubsReply{}
	buf := msgClubReply.GetContent()

	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("err:", err)
	}

	log.Println("reply:", reply)
}

func testLoadClubMembers(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	var url = "http://localhost:3002/lobby/uuid/loadClubMembers?tk=" + tk + "&clubID=83f6aa30-8e63-11e9-8fea-107b445225b6"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	reply := &MsgClubLoadMembersReply{}
	buf := msgClubReply.GetContent()

	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("err:", err)
	}

	log.Println("reply:", reply)
}

func testJoinClub(id string, clubNumber string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	var url = "http://localhost:3002/lobby/uuid/joinClub?tk=" + tk + "&clubNumber=" + clubNumber
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	// reply := &MsgClubLoadMembersReply{}
	// buf := msgClubReply.GetContent()

	// err = proto.Unmarshal(buf, reply)
	// if err != nil {
	// 	log.Println("err:", err)
	// }
	buf := msgClubReply.GetContent()

	log.Println("reply:", string(buf))
}

func testLoadClubEvent(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	var url = "http://localhost:3002/lobby/uuid/loadClubEvents?tk=" + tk + "&clubID=6b512ef0-7b77-11e9-a192-107b445225b6"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	reply := &MsgClubLoadEventsReply{}
	buf := msgClubReply.GetContent()

	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("err:", err)
	}

	log.Println("reply:", reply)
}

func testJoinApproval(id string, applicantID string, clubID string, agree string, eID string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="

	var url = "http://121.196.210.106:30002/lobby/uuid/joinApproval??tk=" + tk +
		"&clubID=" + clubID + "&applicantID=" + applicantID + "&agree=" + agree + "&eID=" + eID
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()

	log.Println("reply:", string(buf))
}

func testClubQuit(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="

	var url = "http://localhost:3002/lobby/uuid/quitClub?tk=" + tk + "&clubID=6b512ef0-7b77-11e9-a192-107b445225b6"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	reply := &MsgClubLoadMyClubsReply{}
	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("err:", err)
	}

	log.Println("reply:", reply)
}

func testCreateClubRoom(id string) {
	tk := lobby.GenTK(id)
	// tk := "vpequ8ELk8xCTPN-heLzghqikggNF85xeH1AyElDSHY="
	config := `{"playerNumAcquired":4, "payNum":0, "payType":0, "handNum":4, "roomType":1, "modName":"game1"}`
	createRoomReq := &lobby.MsgCreateRoomReq{}
	createRoomReq.Config = &config

	buf, err := proto.Marshal(createRoomReq)
	if err != nil {
		log.Println("testCreateClubRoom, error:", err)
		return
	}

	var url = "http://localhost:3002/lobby/uuid/createClubRoom?tk=" + tk + "&clubID=9949cd58-7e97-11e9-a192-107b445225b6"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("OPTIONS", url, bytes.NewBuffer(buf))

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	reply := &lobby.MsgCreateRoomRsp{}
	err = proto.Unmarshal(body, reply)
	if err != nil {
		log.Println("testCreateClubRoom, err:", err)
		return
	}

	if reply.GetResult() != 0 {
		log.Printf("reply errCode:%d, retMsg:%s", reply.GetResult(), reply.GetRetMsg())
	} else {
		log.Println("reply:", reply)
	}
	// log.Println("reply:", reply)
}

func testLoadClubRooms(id string) {
	tk := lobby.GenTK(id)

	var url = "http://localhost:3002/lobby/uuid/loadClubRooms?tk=" + tk + "&clubID=9949cd58-7e97-11e9-a192-107b445225b6"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("POST", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	reply := &lobby.MsgLoadRoomListRsp{}
	err = proto.Unmarshal(body, reply)
	if err != nil {
		log.Println("testCreateClubRoom, err:", err)
		return
	}

	if reply.GetResult() != 0 {
		log.Printf("reply errCode:%d, retMsg:%s", reply.GetResult(), reply.GetRetMsg())
	} else {
		log.Println("reply:", reply)
	}
	// log.Println("reply:", reply)
}

func testDeleteClubRoom(id string) {
	tk := lobby.GenTK(id)

	var url = "http://localhost:3002/lobby/uuid/deleteClubRoom?tk=" + tk + "&clubID=5fab2ce0-7e06-11e9-a192-107b445225b6&roomID=2d4958eb-5162-429d-beb8-0d81509ad89c"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("POST", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	reply := &lobby.MsgDeleteRoomReply{}
	err = proto.Unmarshal(body, reply)
	if err != nil {
		log.Println("testDeleteClubRoom, err:", err)
		return
	}

	if reply.GetResult() != 0 {
		log.Printf("reply errCode:%d", reply.GetResult())
	} else {
		log.Println("reply:", reply)
	}
	// log.Println("reply:", reply)
}

func testLoadMyApplyEvent(id string) {
	tk := lobby.GenTK(id)

	var url = "http://localhost:3002/lobby/uuid/loadMyApplyEvent?tk=" + tk
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	reply := &MsgClubLoadApplyRecordReply{}
	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("err:", err)
	}

	events := reply.GetRecords()
	log.Println("reply:", reply)
	log.Println("event:", len(events))
}

func renameClub(id string, clubID string, clubName string) {
	tk := lobby.GenTK(id)

	var url = "http://localhost:3002/lobby/uuid/renameClub?tk=" + tk + "&clubID=" + clubID + "&clname=" + clubName
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	log.Println("event:", string(buf))
}

func kickOutMember(id string, clubID string, memberID string) {
	// tk := lobby.GenTK(id)

	// var url = "http://localhost:3002/lobby/uuid/kickOut?tk=" + tk + "&clubID=" + clubID + "&memberID=" + memberID
	var url = "https://dfh5-develop.qianz.com:30000/lobby/uuid/kickOut?&tk=jUdigI7QXE_t09t5HdnySKyBJ0_xy1v5fTinYcqKAeU=&clubID=a53f5dee-8e87-11e9-9a1a-00163e0f7404&memberID=10000003"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	log.Println("event:", string(buf))
}

func changeClubMemberRole(id string, clubID string, memberID string, role int) {
	tk := lobby.GenTK(id)

	roleString := fmt.Sprintf("%d", role)
	var url = "http://localhost:3002/lobby/uuid/changeRole?tk=" + tk + "&clubID=" + clubID + "&memberID=" + memberID + "&role=" + roleString
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	log.Println("event:", string(buf))
}

func loadClubManagers(id string, clubID string) {
	tk := lobby.GenTK(id)
	var url = "http://localhost:3002/lobby/uuid/loadClubMgrs?tk=" + tk + "&clubID=" + clubID
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	reply := &MsgClubLoadMembersReply{}
	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("parse error:", err)
	}

	log.Println("reply:", reply)
}


func allowMemberCreateRoom(id string, clubID string, memberID string) {
	tk := lobby.GenTK(id)
	var url = "http://localhost:3002/lobby/uuid/allowMemberCreateRoom?tk=" + tk + "&clubID=" + clubID +"&memberID=" + memberID + "&allowCreateRoom=yes"
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	reply := &MsgClubLoadMembersReply{}
	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("parse error:", err)
	}

	log.Println("reply:", reply)
}

func testSetRoomOptions(id string, clubID string, options string) {
	tk := lobby.GenTK(id)
	var url = "http://localhost:3002/lobby/uuid/setRoomOptions?tk=" + tk + "&clubID=" + clubID + "&options=" + options
	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200, resp.StatusCode:", resp.StatusCode)
		return
	}

	errcode := resp.Header.Get("error")
	if errcode != "" {
		log.Println("errorcode: ", errcode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("handlerChat error:", err)
		return
	}

	msgClubReply := &MsgClubReply{}
	err = proto.Unmarshal(body, msgClubReply)
	if err != nil {
		log.Println("err:", err)
	}

	if msgClubReply.GetReplyCode() == int32(ClubReplyCode_RCError) {
		genericRely := &MsgCubOperGenericReply{}
		err = proto.Unmarshal(msgClubReply.GetContent(), genericRely)
		if err != nil {
			log.Println("parse error:", err)
		}

		log.Println("errCode:", genericRely.GetErrorCode())
		return
	}

	buf := msgClubReply.GetContent()
	if len(buf) == 0 {
		log.Println("len(buf) == 0")
		return
	}

	reply := &MsgClubInfo{}
	err = proto.Unmarshal(buf, reply)
	if err != nil {
		log.Println("parse error:", err)
	}

	log.Println("reply:", reply)
}