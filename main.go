package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hemantgarg925/mymongo/router"
)

func main() {
	fmt.Println("server started")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server is listening at port 4000")
}
