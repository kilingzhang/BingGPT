package websocket

import (
	"github.com/fasthttp/websocket"
	"log"
	"net/http"
)

func Conversation(w http.ResponseWriter, r *http.Request) {
	var upgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
			//allowedOrigins := []string{"https://websocketking.com"}
			//origin := r.Header.Get("Origin")
			//for _, allowed := range allowedOrigins {
			//	if allowed == origin {
			//		return true
			//	}
			//}
			//return false
		},
	}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	handleWebSocket(conn)
}

func handleWebSocket(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("conn.ReadMessage():", err)
			return
		}
		log.Printf("Received message: %s\n", message)

		err = conn.WriteMessage(websocket.TextMessage, []byte("Pong!"))
		if err != nil {
			log.Println("conn.WriteMessage:", err)
			return
		}
	}
}
