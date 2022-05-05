package main

import (
	// "encoding/json"
	"koryos/configs"
	"koryos/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// router := mux.NewRouter()

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
		
	// 	json.NewEncoder(w).Encode(map[string]string{"data": "koryos server is running ..."})
	// }).Methods("GET")

	configs.ConnectDB()

	log.Fatal(http.ListenAndServe(":4000", NewRouter()))
}

func NewRouter() *mux.Router {
    r := mux.NewRouter()
    routes.RoomRoute(r)
	routes.MessageRoute(r)
	routes.DualRoute(r)
    return r
}