package main

import (
	"log"
	"modulo/src/router"
	"net/http"
)

func main() {
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":8080", r))
}
