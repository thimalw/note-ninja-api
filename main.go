package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := routes()

	// TODO: Get port from env var
	fmt.Println("Server listening on port 3080")
	log.Fatal(http.ListenAndServe(":3080", r))
}
