package webserverhelper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ServerConfig : config
type ServerConfig struct {
	IP   string
	Port int
}

var serverConfig ServerConfig

// Start : Starts the server using the specified config
func Start(config ServerConfig) {
	serverConfig = config

	// https://github.com/gorilla/mux#install

	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", bookHandler)
	router.HandleFunc("/", homeHandler)

	http.Handle("/", router)

	server := &http.Server{
		Addr:           config.IP + ":" + strconv.Itoa(config.Port),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Starting server: " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome home")
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("Title: " + vars["title"] + ", page: " + vars["page"])
}
