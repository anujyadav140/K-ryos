package controllers

import (
	"encoding/json"
	"koryos/configs"
	"koryos/models"
	"koryos/responses"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func CreateRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var room models.Room
		//check if the request fits the response
		if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//validate using go validator package
		if validationErr := Validate.Struct(&room); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//insert the request into the database
		createdRoom := configs.DB.Create(&room)
		err := createdRoom.Error

		//if insertion fails
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//success!
		w.WriteHeader(http.StatusCreated)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "Room Created!"}}
		json.NewEncoder(w).Encode(response)

	}
}

func ListRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer timeTrack(time.Now(), "Fetching list rooms info ...")
		var room []models.Room

		//get rooms from the database
		getRooms := configs.DB.Find(&room)
		err := getRooms.Error

		//if get rooms fails
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//get rooms success
		w.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": &room}}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var room models.Room

		findRoom := configs.DB.First(&room, params["id"])

		if findRoom.RowsAffected > 0 {
			configs.DB.Unscoped().Delete(&room)
			w.WriteHeader(http.StatusOK)
			response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": &room}}
			json.NewEncoder(w).Encode(response)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			response := responses.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No user with that ID"}}
			json.NewEncoder(w).Encode(response)
		}
	}
}

