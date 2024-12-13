package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Could not open HTML file", http.StatusInternalServerError)
		return
	}
	defer htmlFile.Close()
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, htmlFile); err != nil {
		http.Error(w, "Failed to send HTML content", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
