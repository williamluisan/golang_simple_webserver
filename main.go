package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadConfig() {
	// load configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// load configuration
	loadConfig()
	port := os.Getenv("WEBSERVER_PORT")

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uri := r.URL.String()
		uriSplit := strings.Split(uri, "/")
		mainPath := uriSplit[1]
		if mainPath != "" {
			http.ServeFile(w, r, mainPath)
		} else {
			http.ServeFile(w, r, "./index.html")
		}
	})

	fmt.Println("Server is running on port " + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}