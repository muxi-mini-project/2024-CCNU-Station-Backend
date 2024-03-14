package chathandler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"guizizhan/model/chat"
	"net/http"
)

func WsHandler(c *gin.Context) {

	uid := c.Query("uid")
	touid := c.Query("touid")
	//chat_type := c.Query("type")

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	//创建一个用户实例
	client := new(chat.Client)
	client = &chat.Client{
		ID:     Createid(uid, touid),
		SendID: Createid(touid, uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	chat.Manager.Register <- client

	//开两个协程用于读写消息
	go Read(client)
	go Write(client)
}

func Createid(uid, touid string) string {
	return uid + "->" + touid
}
