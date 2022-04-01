package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	r := router.Generate()

	fmt.Println("Escutando na porta ", config.ApiPort)
	log.Fatal(http.ListenAndServe(":5000", r))
}
