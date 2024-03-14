package chat

import "time"

type Msg struct {
	SendID    string    `json:"send_id"`
	ReceiveID string    `json:"receive_id"`
	SendTime  time.Time `json:"send_time"`
	Content   string    `json:"content"`
}
