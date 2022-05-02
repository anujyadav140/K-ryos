package controllers

import (
	"koryos/configs"
	"koryos/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var socketMessage string

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		socketMessage = string(p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func WsMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		wsMessage, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)
		}

		newMessage := models.Message{
			Content: socketMessage,
			RoomID:  1,
		}

		if newMessage.Content != "" {
			createdMessage := configs.DB.Create(&newMessage)
			createdMessageErr := createdMessage.Error
			if createdMessageErr != nil {
				log.Println(createdMessageErr)
			}
			socketMessage = ""
		} else {log.Println("EMPTY MESSAGE")}

		reader(wsMessage)
	}
}
