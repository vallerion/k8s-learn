package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	version := os.Getenv("VER_TAG")

	http.HandleFunc("/number", func(writer http.ResponseWriter, request *http.Request) {
		rand.Seed(time.Now().UnixNano())

		_, _ = writer.Write([]byte(strconv.Itoa(rand.Intn(100))))
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		host, _ := os.Hostname()
		_, _ = writer.Write([]byte("Number APP\nVersion: " + version + "\nHost: " + host))
	})

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusAccepted)
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
