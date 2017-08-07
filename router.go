package main

import (
	"github.com/almakinah/gobot/bot"
	"github.com/gorilla/mux"
)

// Initialize the gorilla router to handle requests on different routes
func initRouter() (mx *mux.Router) {
	mx = mux.NewRouter()

	mx.HandleFunc("/makmak/handle{_:/?}", bot.VerifyToken).Methods("GET")
	mx.HandleFunc("/makmak/handle{_:/?}", bot.RecieveMessage).Methods("POST")

	return

}
