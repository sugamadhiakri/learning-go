package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	fmt.Println("Listening and serving at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
