package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func loadConfig() {
	// load configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func webServerHandler(w http.ResponseWriter, r *http.Request) {
	rootFolder := os.Getenv("ROOT_FOLDER")
	filePath := "./" + rootFolder + r.URL.String()
	_, fileName := path.Split(filePath)
	fileNameExt := filepath.Ext(filePath)

	// check request
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// check file name extension
	if fileNameExt == "" {
		// if no extension, default open index.html
		filePath = filePath + "/index.html"
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("file %s not found\n", filePath)
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "404 not found", http.StatusNotFound)
		
		return
	}
	defer file.Close()

	http.ServeContent(w, r, fileName, time.Now(), file)
}

func main() {
	// load configuration
	loadConfig()
	port := os.Getenv("WEBSERVER_PORT")

	http.HandleFunc("/", webServerHandler)

	fmt.Println("Server is running on port " + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}