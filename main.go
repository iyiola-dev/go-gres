package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iyiola-dev/go-gres/router"
)

func main() {
	route := router.Router()
	fmt.Println("starting route at one port like that")
	log.Fatal(http.ListenAndServe(":8080", route))

}
