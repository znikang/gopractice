package gwebsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 根據需要允許特定來源
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if !validateToken(token) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Received:", string(msg))
	}
}

func validateToken(token string) bool {
	token = strings.TrimPrefix(token, "Bearer ")
	return token == "your-valid-token" // 替換成真正的驗證邏輯
}
