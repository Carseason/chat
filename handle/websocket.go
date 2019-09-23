package handle

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type handleMessage struct {
	Id      string
	Code    int
	Message Message
	Ip      string
	Ua      string
	Referer string
}
type Message struct {
	Content string
	Author  Author
	Date    time.Time
}

type Author struct {
	Name   string
	Avatar string
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients    = make(map[string]map[string]*websocket.Conn) // connected clients
	broadcast  = make(chan handleMessage)                    // broadcast channel
	chatRecord = make(map[string][]handleMessage)            // Chat Record map
)

func HandleWebSocket(router *gin.Context) {
	id := router.Query("id")
	if id == "" {
		return
	}
	name := router.DefaultQuery("name", "游客")
	if name == "" {
		return
	}
	avatar := router.DefaultQuery("avatar", "https://i.loli.net/2019/08/06/r7xN8Wq4jOZQ5i3.jpg")
	if avatar == "" {
		avatar = "https://i.loli.net/2019/08/06/r7xN8Wq4jOZQ5i3.jpg"
	}

	ws, err := upgrader.Upgrade(router.Writer, router.Request, nil) //升级get请求为webSocket协议
	if err != nil {
		return
	}
	defer ws.Close()
	if _, ok := clients[id]; !ok {
		clients[id] = make(map[string]*websocket.Conn)
	}
	clients[id][name] = ws
	for {
		code, message, err := ws.ReadMessage() // 读取ws中的数据
		if err != nil {
			delete(clients[id], name)
			break
		}
		content := string(message)
		if content == "ping" {
			//ws.WriteMessage(code, []byte("pong"))
			//ws.WriteMessage(mt, message) //二进制返回消息
			//	 ws.WriteJSON(gin.H{"message": "返回的消息"})	//JSO返回消息
			ws.WriteJSON(gin.H{
				"code":    code,
				"message": "pong",
				"data":    recordToMessage(getRecord(id)),
			})

		} else {
			content = replaceEmoji(content)
			content = regexpReplaceAllString(content, `(独立|台独|港独|你妈|革命|妈逼|操逼)`, "**")
			messageValue := handleMessage{
				Id:      id,
				Code:    code,
				Ip:      router.ClientIP(),
				Ua:      router.GetHeader("User-Agent"),
				Referer: router.GetHeader("Referer"),
				Message: Message{
					Content: content,
					Date:    time.Now(),
					Author: Author{
						Name:   name,
						Avatar: avatar,
					},
				},
			}
			setRecord(id, messageValue)
			broadcast <- messageValue
		}

	}
	delete(clients[id], name)
}
func HandleMessages() {
	for {
		select {
		case msg := <-broadcast:
			for value := range clients[msg.Id] {
				if err := clients[msg.Id][value].WriteJSON(gin.H{
					"code":    msg.Code,
					"message": nil,
					"data":    msg.Message,
				}); err != nil {
					delete(clients[msg.Id], value)
					break
				}
			}
		}
	}
}

/****************聊天记录*******************/
func recordToMessage(value []handleMessage) []Message {
	message := []Message{}
	for v := range value {
		message = append(message, value[v].Message)
	}
	return message
}

func getRecord(id string) []handleMessage {
	return chatRecord[id]
}
func setRecord(id string, value handleMessage) {
	recordArray := chatRecord[id]
	t := len(recordArray)
	if t > 100 {
		recordArray = recordArray[1:]
	}
	recordArray = append(recordArray, value)
	chatRecord[id] = recordArray
}
