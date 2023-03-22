package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/papihack/mongoapi/router"
)

const PORT = 4000

func main() {
	fmt.Println("[*] MongoDB API")
	r := router.Router()
	fmt.Println("[*] Server is getting started....")
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT), r))
	fmt.Println("[*] Listening on port", strconv.Itoa(PORT) + "...")
}
