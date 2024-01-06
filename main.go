package main

import (
	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/driver"
	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/handler"
	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/service"
	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/store"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main() {
	db := driver.ConnectToMongoDb()
	st := store.NewMongoUserStore(db)
	svc := service.NewAuthService(st)
	list := handler.NewUserHandler(*svc)

	r := mux.NewRouter()

	r.HandleFunc("/register", list.RegistrationHandler).Methods("POST")

    http.Handle("/", r)

    fmt.Println("Server is running on :8080...")
    http.ListenAndServe(":8080", nil)
}
