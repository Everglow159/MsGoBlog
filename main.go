package main

import (
	"log"
	"msgoblog/route"
	"net/http"
)

func init() {

}

////////////////////////MAIN//////////////////////////////
func main() {
	//only one main for one project
	//This is web project demo
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	route.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
