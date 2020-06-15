package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//当前已连接用户集合
	connectionUser = make(map[string]*websocket.Conn)
	//接收消息通道 0:发消息连接fd 1:消息主体
	wsChannel = make(chan [][]byte)
	//关闭消息通知通道
	wsChannelCancel = make(chan struct{})
)

func init() {
	//监听消息并群发的协程,发送给除自己外的所有用户（如果存在的话）
	go func() {
		for {
			select {
			case msg := <-wsChannel:
				fmt.Println(string(msg[1]))
				//遍历用户连接集合
				if len(connectionUser) < 2 {
					fmt.Println("无其他用户连接，无需推送")
					continue
				}
				for _, v := range connectionUser {
					//发消息本人不需要推送消息
					fmt.Println("发消息用户：" + string(msg[0]))
					fmt.Println("接收消息用户：" + v.RemoteAddr().String())
					if v.RemoteAddr().String() == string(msg[0]) {
						fmt.Println("发送消息用户，无需推送")
						continue
					}
					if err := v.WriteMessage(1, msg[1]); err != nil {
						fmt.Println("发送消息失败，错误原因：" + err.Error())
						//发送消息失败，断开连接，并从登陆集合中清除连接用户
						v.Close()
						//获取退出用户的fd，从用户集合中删除
						delete(connectionUser, v.RemoteAddr().String())
						fmt.Println("用户" + v.RemoteAddr().String() + "退出")
						fmt.Println("剩余用户数量还有：" + strconv.Itoa(len(connectionUser)) + "人")
					}
				}
			case <-wsChannelCancel:
				//ws主线程退出 跳出for循环
				fmt.Println("main wsProcess exit")
				goto EXIT
			}
		}
	EXIT:
		fmt.Println("bye")
	}()
}

func WebsocketStart(w http.ResponseWriter, r *http.Request) {
	conn, ok := upgrader.Upgrade(w, r, nil) // 实际应用时记得做错误处理
	if ok != nil {
		fmt.Println(ok.Error())
		return
	}
	//连接退出/关闭
	defer connectCancel(conn)
	//提示连接成功
	conn.WriteMessage(1, []byte(`{"msg":"welcome to websocket!!"}`))
	//添加到用户集合通道
	connectionUser[conn.RemoteAddr().String()] = conn
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println("当前连接用户数量有：" + strconv.Itoa(len(connectionUser)) + "人")
	for {
		// 读取客户端的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("err msg:" + err.Error())
			goto Err
		}
		//消息信息放进消息通道
		message := [][]byte{0: []byte(conn.RemoteAddr().String()), 1: msg}
		wsChannel <- message
	}
Err:
	//读取消息失败，断开连接，并从登陆集合中清除连接用户
	conn.Close()
	//获取退出用户的fd，从用户集合中删除
	delete(connectionUser, conn.RemoteAddr().String())
	fmt.Println("用户" + conn.RemoteAddr().String() + "退出")
	fmt.Println("剩余用户数量还有：" + strconv.Itoa(len(connectionUser)) + "人")
}

func connectCancel(conn *websocket.Conn) {
	conn.Close()
	//防止可能出现的内存泄漏问题，需要人为关闭相应通道
	//close(wsChannelCancel)
	//close(wsChannel)
}
