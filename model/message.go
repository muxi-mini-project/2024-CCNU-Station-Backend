package model

//
//import (
//	"fmt"
//	"github.com/gorilla/websocket"
//	"gopkg.in/fatih/set.v0"
//	"gorm.io/gorm"
//	"net"
//	"net/http"
//)
//
//type Message struct {
//	gorm.model
//	FromID   string //发送者
//	TargetID string //接收者
//	Type     string //发送类型
//	Media    int    //消息类型
//	Content  string //消息内容
//	Pic      string
//	Url      string
//	Desc     string
//	Amount   int //大小
//}
//type Node struct {
//	Conn      *websocket.Conn
//	DataQueue chan []byte
//	GroupSets set.Interface
//}
//
//func Chat(writer http.ResponseWriter, request *http.Request, isvalida bool) {
//	query := request.URL.Query()
//	senderid := query.Get("senderid")
//	targetid := query.Get("targetid")
//	context := query.Get("context")
//	msgtype := query.Get("type")
//	conn, err := (&websocket.Upgrader{
//		CheckOrigin: func(r *http.Request) bool {
//			return isvalida
//		},
//	}).Upgrade(writer, request, nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//获取conn
//	node := &Node{
//		Conn:      conn,
//		DataQueue: make(chan []byte, 50),
//		GroupSets: set.New(set.ThreadSafe),
//	}
//	//用户关系
//
//	//将senderid和node绑定并加锁
//	rwLocker.Lock()
//	clientMap[senderid] = node
//	rwLocker.Unlock()
//
//	//完成发送逻辑
//	go sendproc(node)
//	//完成接收逻辑
//	go recvproc(node)
//}
//func sendproc(node *Node) {
//	for {
//		select {
//		case data := <-node.DataQueue:
//			err := node.Conn.WriteMessage(websocket.TextMessage, data)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//
//	}
//}
//
//func recvproc(node *Node) {
//	for {
//		_, data, err := node.Conn.ReadMessage()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		broadMsg(data)
//		fmt.Println("[ws]<<<<", data)
//
//	}
//}
//
//var udpsendChan chan []byte = make(chan []byte, 1024)
//
//func broadMsg(data []byte) {
//	udpsendChan <- data
//}
//func init() {
//	go udpSendProc()
//	go udpRecvProc()
//}
//
//// 完成udp数据发送协程
//func udpSendProc() {
//	net.DialUDP("udp",nil,&net.UDPAddr{
//		IP: net.IPv4(172.30.144.1),
//		Port: 8080
//	})
//	defer con.Close()
//}
