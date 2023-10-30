package main

import (
	"fmt"
	"log"
	"mydiary/config"
	"mydiary/internal/api/handlers"
	"mydiary/internal/database"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init () {
	config.LoadEnvVariables()
	// db connect
	database.ConnectDB()
}


func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/healthcheck", healthcheck)
	router.HandleFunc("/api/auth/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/api/auth/login", handlers.LoginHandler).Methods("POST")

	//var dir string = "./static"
	// This will serve files under http://localhost:8000/static/<filename>
    //router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

    srv := &http.Server{
        Handler:      router,
        Addr: "localhost:8000",	
        //enforced timeouts
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	fmt.Println("Server is running on localhost:8000")
	log.Fatal(srv.ListenAndServe())
}



func healthcheck (w http.ResponseWriter, r *http.Request) {

	fmt.Println("it is healthy")
}
