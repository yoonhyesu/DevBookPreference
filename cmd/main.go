package main

import (
	"DBP/internal/router"
	"DBP/pkg/database"
	"log"
)

func main() {
	client := database.NewMariaDBConnection()
	defer client.Close()

	router := router.SetupRouter(client)
	log.Fatal(router.Run(":7031"))
}
