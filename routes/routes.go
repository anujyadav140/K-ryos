package routes

import (
	"koryos/controllers"

	"github.com/gorilla/mux"
)

func RoomRoute(router *mux.Router) {
	router.HandleFunc("/room", controllers.CreateRoom()).Methods("POST")
	router.HandleFunc("/list-rooms", controllers.ListRooms()).Methods("GET")
	router.HandleFunc("/rooms/{id}", controllers.DeleteRoom()).Methods("DELETE")
}

func MessageRoute(router *mux.Router) {
	router.HandleFunc("/room/message", controllers.CreateMessage()).Methods("POST")
	router.HandleFunc("/room/ws-message", controllers.WsMessage())
	router.HandleFunc("/room/messages", controllers.GetMessages()).Methods("GET")
	router.HandleFunc("/room/messages/{RoomID}", controllers.GetMessagesByID()).Methods("GET")
}