package common

const (
	SUCCESS               = 1000 //成功
	FAIL                  = 1001 //失败
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400
	ErrorDatabase         = 40001

	WebsocketSuccessMessage = 50001
	WebsocketSuccess        = 50002 //websocket成功注册
	WebsocketEnd            = 50003
	WebsocketOnlineReply    = 50004 //对方在线应答
	WebsocketOfflineReply   = 50005 //对方不在线应答
	WebsocketLimit          = 50006
)
