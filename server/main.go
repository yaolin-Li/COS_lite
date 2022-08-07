package main

import (
	"log"
	"net/http"
	"os"
	"github.com/COS_lite/server/objects"
)

func main() {
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}