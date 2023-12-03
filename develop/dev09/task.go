package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	get("https://pkg.go.dev/net/http#example-Get")
}

func get(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	file.Write(body)
	file.Close()
}
