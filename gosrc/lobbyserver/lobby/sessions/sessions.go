package sessions

import (
	"lobbyserver/config"
	"lobbyserver/lobby"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"

	"github.com/gorilla/mux"
)

const (
	wsReadLimit       = 64 * 1024 // 64K
	wsReadBufferSize  = 4 * 1024
	wsWriteBufferSize = 4 * 1024
)

var (
	upgrader = websocket.Upgrader{ReadBufferSize: wsReadBufferSize, WriteBufferSize: wsWriteBufferSize}

	userMgr *UserMgr
)

func waitWebsocketMessage(ws *websocket.Conn, user *User, r *http.Request) {
	log.Printf("wait ws msg, userId: %s, peer: %s", user.userID(), r.RemoteAddr)

	ws.SetPongHandler(func(msg string) error {
		user.lastReceivedTime = time.Now()
		return nil
	})

	ws.SetPingHandler(func(msg string) error {
		user.lastReceivedTime = time.Now()
		user.sendPong(msg)
		return nil
	})

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(" websocket receive error:", err)
			ws.Close()
			user.onWebsocketClosed(ws)
			break
		}

		user.lastReceivedTime = time.Now()

		if message != nil && len(message) > 0 {
			user.onWebsocketMessage(ws, message)
		}

		//log.Printf("receive from user %d message:%s", user.userID(), message)
	}
	log.Printf("ws closed, userId %s, peer:%s", user.userID(), r.RemoteAddr)
}

func tryAcceptUser(ws *websocket.Conn, r *http.Request) {
	userID, ok := lobby.VerifyToken(r)
	if !ok {
		log.Println("verifyUser failed")
		lobby.ReplyLoginError(ws, int32(lobby.LoginState_ParseTokenError))
		return
	}

	var user = newUser(ws, userID)

	oldUser := userMgr.getUserByID(user.uID)
	if oldUser != nil {
		oldUser.detach()
		oldUser.wg.Wait()
	}

	userMgr.addUser(user)

	lobby.LoginReply(ws, userID)

	defer func() {
		userMgr.removeUser(user)
		user.wg.Done()
	}()

	waitWebsocketMessage(ws, user, r)
}

func acceptWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 接收限制
	ws.SetReadLimit(wsReadLimit)

	defer ws.Close()

	log.Println("accept websocket:", r.URL)
	tryAcceptUser(ws, r)
}

// InitWith init
func InitWith(mainRouter *mux.Router) {
	userMgr = newUserMgr()

	startAliveKeeper()

	lobby.SessionMgr = userMgr

	var sessions = mainRouter.PathPrefix("/ws").Subrouter()
	sessions.HandleFunc("/", acceptWebsocket)

	loadSensitiveWordDictionary(config.SensitiveWordFilePath)
}