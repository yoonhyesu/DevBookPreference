package main

import (
	"SpaceDev/pkg/database"
	"log"
	"net/http"
)

func main() {
	client := database.NewMariaDBConnection()
	defer client.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
