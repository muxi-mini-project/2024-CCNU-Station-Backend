package chathandler

import (
	"fmt"
	"gorm.io/gorm"
	"guizizhan/model/chat"
	"time"
)

func StoreMsg(db *gorm.DB, id string, msg string) {
	var sendid, recipientid string
	fmt.Sscanf(id, "%s->%s", &sendid, &recipientid)
	var chatmsg = chat.ChatMessage{
		SendID:      sendid,
		RecipientID: recipientid,
		Content:     msg,
		SendTime:    time.Now(),
	}
	db.Create(&chatmsg)
}
