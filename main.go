package main

import (
	"fmt"
	"log"
	"net/http"
	"stocks/main/router"
)

func main(){

	r:=router.Router()
	fmt.Println("Server listening on 127.0.0.1:3000")

	log.Fatal(http.ListenAndServe(":3000",r))
}