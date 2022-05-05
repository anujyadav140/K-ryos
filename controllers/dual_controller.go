package controllers

import (
	"encoding/json"
	"koryos/configs"
	"koryos/models"
	"koryos/responses"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var wg sync.WaitGroup
var rooms *[]models.Room
var messages *[]models.Message

func DualController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer timeTrack(time.Now(), "Fetching dual controller info ...")

		wg.Add(2)
		params := mux.Vars(r)

		go GetRooms()

		go GetMessage(params["RoomID"])

		wg.Wait()

		w.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"rooms": rooms,"messages": messages}}
		json.NewEncoder(w).Encode(response)
	}
}

func GetRooms() {
	defer wg.Done()
	var room []models.Room

	getRooms := configs.DB.Find(&room)
	err := getRooms.Error

	if err != nil {
		log.Println("CANNOT RETRIEVE ROOM DATA!")
	}

	rooms = &room
	// i := 0
	// for i < len(room){
	// 	rooms = append(rooms, room[i])
	// 	i++
	// }
}

func GetMessage(iden string) {
	defer wg.Done()
	var message []models.Message

	findMessages := configs.DB.Find(&message, "room_id = ?", iden)
	err := findMessages.Error

	if err != nil {
		log.Println("CANNOT RETRIEVE MESSAGE DATA!")
	}

	messages = &message

	// j := 0
	// for j < len(message){
	// 	messages = append(messages, message[j])
	// 	j++
	// }
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
   }