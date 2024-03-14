package chat

import (
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"time"
)

type ChatMessage struct {
	gorm.Model
	SendID      string    //发送者的ID
	RecipientID string    //接收者的ID
	GroupID     int       //群ID，该消息要发到哪个群里面
	Content     string    //内容
	SendTime    time.Time //发送的时间
}

type Group struct {
	ID        int    `gorm:"primaryKey"` //群ID
	Groupname string `json:"group_name"`
}
type SendMsg struct {
	Type        int    `json:"type"`
	SendID      string `json:"send_id"`
	RecipientID string `json:"recipient_id"`
	Content     string `json:"content"`
}
type ReplyMsg struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
	From    string `json:"from"`
}
type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan *Broadcast),
	Reply:      make(chan *Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}
