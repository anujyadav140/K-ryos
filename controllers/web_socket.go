package controllers

import (
	"encoding/json"
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

type socketPayload struct {
	Content string
	Room int
}

var payload socketPayload

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		} 
		
		if err := json.Unmarshal(p, &payload); err != nil {
			log.Println("failed to unmarshal:", err)
		}
		log.Printf("Content: %s, Room_id: %b", payload.Content, payload.Room)

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
			Content: payload.Content,
			RoomID:  payload.Room,
		}

		if newMessage.Content != "" {
			createdMessage := configs.DB.Create(&newMessage)
			createdMessageErr := createdMessage.Error
			if createdMessageErr != nil {
				log.Println(createdMessageErr)
			}
			newMessage.Content = ""
		} else {log.Println("EMPTY MESSAGE")}

		reader(wsMessage)
	}
}
