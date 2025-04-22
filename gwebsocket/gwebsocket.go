package gwebsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"sync"
	"time"
)

var SocketClients = make(map[string]*websocket.Conn)
var clientsLock = sync.RWMutex{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 根據需要允許特定來源
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	userID := validateAndGetUserID(token)

	if !validateToken(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	if userID == "" {
		conn.WriteMessage(websocket.TextMessage, []byte("unauthorized"))
		return
	}
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	// 設定 pong handler，收到 pong 時延長存活
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	clientsLock.Lock()
	SocketClients[userID] = conn
	clientsLock.Unlock()

	//defer conn.Close()

	defer func() {
		clientsLock.Lock()
		delete(SocketClients, userID)
		clientsLock.Unlock()
		conn.Close()
		fmt.Println("User disconnected:", userID)
	}()

	// 在 goroutine 裡定時送 ping
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			err := conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				fmt.Println("Ping error, closing connection")
				conn.Close()

				clientsLock.Lock()
				delete(SocketClients, userID)
				clientsLock.Unlock()

				return
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Received:", string(msg))
	}
}

// 主動發送訊息給特定使用者
func sendMessageToUser(userID, message string) {
	clientsLock.RLock()
	conn, ok := SocketClients[userID]
	clientsLock.RUnlock()
	if ok {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Printf("Send to %s failed: %v\n", userID, err)
		}
	} else {
		fmt.Printf("User %s not connected\n", userID)
	}
}

func validateAndGetUserID(token string) string {
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "token-user1" {
		return "user1"
	}
	if token == "token-user2" {
		return "user2"
	}
	return ""
}

func validateToken(token string) bool {
	token = strings.TrimPrefix(token, "Bearer ")
	return token == "your-valid-token" // 替換成真正的驗證邏輯
}
