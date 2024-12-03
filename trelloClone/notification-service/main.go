package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Povezivanje sa bazama
	ConnectDatabase()
	defer CloseDatabase()
	ConnectUsersService()

	// Postavljanje ruta
	r := mux.NewRouter()
	r.HandleFunc("/notifications", GetNotificationsHandler).Methods("GET")

	//log.Println("Notification Service is running on port 8081")
	//log.Fatal(http.ListenAndServe(":8081", r))

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Dozvoli sve origin-e; za specifične origin-e, koristi {"http://localhost:4200"}

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(r)))
}
