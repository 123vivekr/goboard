package main

import (
	"log"
	"net/http"
)

func main() {
	server := &GoBoardServer{NewBoltScoreStore()}
	log.Fatal(http.ListenAndServe(":6000", server))
}
