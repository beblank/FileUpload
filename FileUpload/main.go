package main

import (
	"FileUpload/file"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", file.Upload)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
