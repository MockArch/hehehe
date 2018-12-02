package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/me/example/app/route"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Applicatio has been started on : 8080")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("Site")
	log.Info(key)

	//creating instance of mux
	routes := route.GetRoutes()
	http.Handle("/", routes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
