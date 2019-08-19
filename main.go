package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"

	log "github.com/sirupsen/logrus"
)

type config struct {
	Host string `default:":"`
	Port int    `default:"8080"`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	var c config
	err := envconfig.Process("DEMO", &c)

	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	http.Handle("/", r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    fmt.Sprintf("%s%d", c.Host, c.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info(fmt.Sprintf("Server listening on: %s%d", c.Host, c.Port))
	log.Fatal(srv.ListenAndServe())
}
