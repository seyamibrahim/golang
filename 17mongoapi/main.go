package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seyamibrahim/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8000",r))
}
