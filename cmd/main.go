package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SaravananSai07/swen"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	customSearchEngineID := os.Getenv("CUSTOM_SEARCH_ENGINE_ID")
	service := swen.NewService(apiKey, customSearchEngineID)
	controller := swen.NewController(service)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Its all good man!"))
	})

	http.HandleFunc("/news", controller.GetNews)

	port := os.Getenv("PORT")
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
