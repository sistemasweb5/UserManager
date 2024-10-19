package main

import (
	endpoint "api/testing/endpoints"
	"log"
)

func main(){
    log.SetFlags( log.Lshortfile)
    // endpoint.InitCognito("localhost:8181")

    log.Print("Starting API testing...")
    restApiAddress := "localhost:5200"
    endpoint.InitClient(&restApiAddress)
}

