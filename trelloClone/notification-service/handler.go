package main

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectUsersService() {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	mongoClient = client
	log.Println("Connected to Users Service MongoDB")
}

func GetNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	// Dohvatanje svih notifikacija
	var notifications []Notification
	iter := cassandraSession.Query("SELECT id, user_id, message, created_at FROM notifications").Iter()

	var notif Notification
	for iter.Scan(&notif.ID, &notif.UserID, &notif.Message, &notif.CreatedAt) {
		notifications = append(notifications, notif)
	}

	if err := iter.Close(); err != nil {
		http.Error(w, "Failed to fetch notifications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}
