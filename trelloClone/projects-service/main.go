package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/projects", GetProjectsHandler).Methods("GET")
	r.HandleFunc("/projects", CreateProjectHandler).Methods("POST")
	r.HandleFunc("/projects/{id}/add-member", AddMemberHandler).Methods("POST")
	r.HandleFunc("/projects/{id}/remove-member", RemoveMemberHandler).Methods("POST")

	//log.Println("Projects Service is running on port 8080...")
	//log.Fatal(http.ListenAndServe(":8080", r))

	//handler := cors.Default().Handler(r)
	//
	//http.Handle("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Dozvoli sve origin-e; za specifične origin-e, koristi {"http://localhost:4200"}

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))

}
