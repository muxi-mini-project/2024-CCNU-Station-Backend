package chathandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"guizizhan/common"
	"guizizhan/model/chat"
)

func Hub(db *gorm.DB) {
	for {
		fmt.Println("<-----监听通信管道----->")
		select {
		case conn := <-chat.Manager.Register:
			replymsg := chat.ReplyMsg{
				Code:    common.WebsocketSuccess,
				Content: "已连接至服务器",
				From:    "SYSTEM",
			}
			msg, _ := json.Marshal(replymsg)
			conn.Socket.WriteMessage(websocket.TextMessage, msg)
			chat.Manager.Clients[conn.ID] = conn
		case conn := <-chat.Manager.Unregister:
			replymsg := chat.ReplyMsg{
				Code:    common.WebsocketEnd,
				From:    "SYSTEM",
				Content: "连接断开",
			}
			msg, _ := json.Marshal(replymsg)
			conn.Socket.WriteMessage(websocket.TextMessage, msg)
			delete(chat.Manager.Clients, conn.ID)
		case broadcast := <-chat.Manager.Broadcast:
			message := broadcast.Message
			id := broadcast.Client.ID
			recipientID := broadcast.Client.SendID
			flag := false
			for _, conn := range chat.Manager.Clients {
				if conn.ID == recipientID {
					select {
					case conn.Send <- message:
						flag = true
					default:
						close(conn.Send)
						delete(chat.Manager.Clients, conn.ID)
					}
				}
			}
			if flag {
				repltmsg := chat.ReplyMsg{
					Code:    common.WebsocketOnlineReply,
					From:    "SYSTEM",
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(repltmsg)
				broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				StoreMsg(db, id, fmt.Sprintf("%s", string(message)))
			} else {
				replymsg := chat.ReplyMsg{
					Code:    common.WebsocketOfflineReply,
					From:    "SYSTEM",
					Content: "对方不在线应答",
				}
				msg, _ := json.Marshal(replymsg)
				broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				StoreMsg(db, id, fmt.Sprintf("%s", string(message)))
			}

		}
	}
}
