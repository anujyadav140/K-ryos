package controllers

import (
	"encoding/json"
	"koryos/configs"
	"koryos/models"
	"koryos/responses"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var Validate = validator.New()

func CreateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message models.Message
		//check if the request fits the response
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//validate using go validator package
		if validationErr := Validate.Struct(&message); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//insert the request into the database
		createdRoom := configs.DB.Create(&message)
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
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "Message sent!"}}
		json.NewEncoder(w).Encode(response)
	}
}

func GetMessages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var messages []models.Message

		getMessages := configs.DB.Find(&messages)
		err := getMessages.Error

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//get messages success
		w.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": &messages}}
		json.NewEncoder(w).Encode(response)
	}
}

func GetMessagesByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var messages []models.Message

		findMessages := configs.DB.Find(&messages, params["RoomID"])
		err := findMessages.Error

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//get messages success
		w.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": &messages}}
		json.NewEncoder(w).Encode(response)
	}
}