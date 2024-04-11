package main

import (
	"fmt"
	"log"
	"login-jwt-example/pkg/api"
	"login-jwt-example/pkg/db"
	"net/http"
)

func main() {
	log.Print("server has started")
    //start the db
    pgdb, err := db.StartDB()
    if err != nil {
        log.Printf("error starting the database %v", err)
    }
    //get the router of the API by passing the db
    router := api.StartAPI(pgdb)
    //get the port from the environment variable
    port := "4001"//os.Getenv("PORT")
    //pass the router and start listening with the server
    err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
    if err != nil {
        log.Printf("error from router %v\n", err)
    }
}