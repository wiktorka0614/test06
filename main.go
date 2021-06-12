package main

import (
	"fmt"
	"log"
	"net/http"
	"test/handler"
)

func main() {
	server := handler.Handler()
	fmt.Println("your server is running...")
	if err := http.ListenAndServe(":3001", server); err != nil {
		log.Fatalln("sth went wrong :(")
	}
}
