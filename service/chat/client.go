package chat

import "github.com/gorilla/websocket"

type User struct {
	ID   string `json:"id"`
	Conn *websocket.Conn
}
