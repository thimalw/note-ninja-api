package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thimalw/note-ninja-api/internal/config"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Panicln("Configuration error.", err)
	}

	r := routes(conf.Database)

	// TODO: Get port from env var
	fmt.Println("Server listening on port: " + conf.Constants.PORT)
	log.Fatal(http.ListenAndServe(":"+conf.Constants.PORT, r))
}
