package chathandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"guizizhan/common"
	"guizizhan/model/chat"
)

func Read(client *chat.Client) {
	defer func() { // 避免忘记关闭，所以要加上close
		chat.Manager.Unregister <- client
		_ = client.Socket.Close()
	}()
	for {
		client.Socket.PongHandler()
		sendmsg := new(chat.SendMsg)
		err := client.Socket.ReadJSON(&sendmsg)
		if err != nil {
			fmt.Println("数据格式不正确")
			chat.Manager.Unregister <- client
			client.Socket.Close()
			break
		}
		if sendmsg.Type == 1 { //	发送单聊消息
			chat.Manager.Broadcast <- &chat.Broadcast{
				Client:  client,
				Message: []byte(sendmsg.Content),
				Type:    1,
			}
			replymsg := chat.ReplyMsg{
				Code:    common.WebsocketSuccessMessage,
				From:    "STSTEM",
				Content: "消息成功发送",
			}

			msg, _ := json.Marshal(replymsg)
			client.Socket.WriteMessage(websocket.TextMessage, msg)

		}
	}
}

func Write(client *chat.Client) {
	defer func() {
		_ = client.Socket.Close()
	}()
	for {
		select {
		//读取管道里面的消息
		case message, ok := <-client.Send:
			//连接不到就返回消息
			if !ok {
				_ = client.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replymsg := chat.ReplyMsg{
				From:    client.SendID,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replymsg)
			_ = client.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
