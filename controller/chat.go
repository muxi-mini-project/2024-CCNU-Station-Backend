package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"guizizhan/service/chat"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[string]*chat.User)

// PrivateChat 私聊
// @Summary 进行WebSocket私聊
// @Description 先将连接升级为WebSocket，以便进行私聊。
// @Accept json
// @Produce json
// @Param uid query string true "用户ID"
// @Param message body chat.Msg true "私聊消息"
// @Api(tags="聊天")
// @Router /api/talk/private_chat [get]
func PrivateChat(user *chat.User) {
	// 将 User 结构加入到 clients 映射中
	clients[user.ID] = user
	//获取websocket连接
	for {
		var msg chat.Msg
		err := user.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			delete(clients, user.ID)
			break
		}
		if msg.ReceiveID != "" {
			for clientID, client := range clients {
				if clientID == msg.ReceiveID {
					err := client.Conn.WriteJSON(msg)
					if err != nil {
						fmt.Printf("error: %v\n", err)
						client.Conn.Close()
						delete(clients, clientID)
					}
					break
				}
			}
		} else {
			break
		}
	}
}

// PublicChat 群聊
// @Summary 进行WebSocket群聊
// @Description 先将连接升级为WebSocket，以便进行群聊。
// @Accept json
// @Produce json
// @Param uid query string true "用户ID"
// @Param message body chat.Msg true "群聊消息"
// @Api(tags="聊天")
// @Router /api/talk/public_chat [get]
func PublicChat(user *chat.User) {
	// 将 User 结构加入到 clients 映射中
	clients[user.ID] = user
	//获取websocket连接
	for {
		var msg chat.Msg
		err := user.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			delete(clients, user.ID)
			break
		}
		if msg.ReceiveID != "" {
			for clientID, client := range clients {
				if clientID == msg.ReceiveID {
					err := client.Conn.WriteJSON(msg)
					if err != nil {
						fmt.Printf("error: %v\n", err)
						client.Conn.Close()
						delete(clients, clientID)
					}
					break
				}
			}
		} else {
			break
		}
	}
}

func Handelewebsocket(c *gin.Context) *websocket.Conn {
	//获取websocket连接
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		panic(err)
	}
	return ws
}
func Createclient(c *gin.Context) chat.User {
	uid := c.Query("uid")
	conn := Handelewebsocket(c)
	var user = chat.User{ID: uid, Conn: conn}
	return user
}
