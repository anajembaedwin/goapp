package main

import (
	"database/sql"
	"goapp/api/handler"
	"goapp/api/middleware"
	"goapp/pkg/repository"
	"goapp/pkg/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
    // Connect to the database.
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=goapp sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

	// Initialize the repository, service, and handler.
    userRepository := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepository)
    userHandler := handler.NewUserHandler(userService)

    // Setup the router and routes.
    router := mux.NewRouter()

	// Apply the middleware to the router
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthenticationMiddleware)

	// user routes
    router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
    router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

    // Start the server.
    log.Fatal(http.ListenAndServe(":8080", router))
}