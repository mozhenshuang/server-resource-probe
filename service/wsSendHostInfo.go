/**
 * @Author root$
 * @Date 2023/3/26$
 * @Note 使用websocket定时推送服务器硬件资源信息
 **/

package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

//Conn类型表示WebSocket连接。服务器应用程序从HTTP请求处理程序调用Upgrader.Upgrade方法以获取* Conn：
var (
	upgrader = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 10,
		// 写入存储空间大小
		WriteBufferSize: 10,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//   完成握手 升级为 WebSocket长连接，使用conn发送和接收消息。
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	//调用连接的WriteMessage和ReadMessage方法以一片字节发送和接收消息。实现如何回显消息：
	//p是一个[]字节，messageType是一个值为websocket.BinaryMessage或websocket.TextMessage的int。
	for {
		hostByteArr, _ := json.Marshal(HostInfo())
		c := newWithSeconds()
		spec := "*/1 * * * * ?"
		c.AddFunc(spec, func() {
			if err := conn.WriteMessage(1, hostByteArr); err != nil {

				log.Println("Writeing error...", err)
				c.Stop()
				return
			}
			log.Printf("Write msg to client: recved: %s \n", hostByteArr)
		})
		c.Start()
		select {}
	}
}

//
// StartWS
//  @Description: 启动websocket服务
//
func StartWS(port string) {
	log.Print("ws linsten on ", port)
	http.HandleFunc("/", wsHandler) // ws://127.0.0.1:8888/
	// 监听 地址 端口
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}
