package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	// Validation
	project.ID = primitive.NewObjectID()
	project.Members = []string{} // Inicijalizujte members kao prazan niz
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	collection := db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func AddMemberHandler(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	var req struct {
		MemberID  string `json:"memberId"`
		ManagerID string `json:"managerId"`
	}

	// Log request payload
	log.Printf("AddMemberHandler called with projectID: %s", projectID)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	collection := db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	projectObjectID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		log.Printf("Invalid project ID: %v", err)
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var project Project
	if err := collection.FindOne(ctx, bson.M{"_id": projectObjectID}).Decode(&project); err != nil {
		log.Printf("Project not found: %v", err)
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Log project data
	log.Printf("Project found: %+v", project)

	// Validations
	if project.ManagerID != req.ManagerID {
		log.Printf("Unauthorized access by ManagerID: %s", req.ManagerID)
		http.Error(w, "You are not authorized to manage this project", http.StatusForbidden)
		return
	}

	if project.EndDate.Before(time.Now()) {
		log.Printf("Cannot add members to a completed project")
		http.Error(w, "Cannot add members to a completed project", http.StatusBadRequest)
		return
	}

	if len(project.Members) >= project.MaxMembers {
		log.Printf("Project is at maximum capacity")
		http.Error(w, "Project is at maximum capacity", http.StatusBadRequest)
		return
	}

	for _, member := range project.Members {
		if member == req.MemberID {
			log.Printf("Member already added: %s", req.MemberID)
			http.Error(w, "Member already added", http.StatusBadRequest)
			return
		}
	}

	// Update project
	update := bson.M{"$push": bson.M{"members": req.MemberID}, "$set": bson.M{"updated_at": time.Now()}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": projectObjectID}, update)
	if err != nil {
		log.Printf("Failed to update project: %v", err)
		http.Error(w, "Failed to add member", http.StatusInternalServerError)
		return
	}

	// Success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Member added successfully")
}

func RemoveMemberHandler(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	var req struct {
		MemberID  string `json:"memberId"`
		ManagerID string `json:"managerId"` // ID menadžera koji vrši zahtev
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	collection := db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	projectObjectID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var project Project
	if err := collection.FindOne(ctx, bson.M{"_id": projectObjectID}).Decode(&project); err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	// Proverite da li je zahtev poslat od menadžera projekta
	if project.ManagerID != req.ManagerID {
		http.Error(w, "You are not authorized to manage this project", http.StatusForbidden)
		return
	}

	// Proverite da li član nije dodeljen zadacima u izradi (simulacija)
	// TODO: Implement real task validation logic
	// Primer: var tasks []Task => Proverite task.status != "In Progress"

	// Proverite da li je član deo projekta
	memberExists := false
	for _, member := range project.Members {
		if member == req.MemberID {
			memberExists = true
			break
		}
	}
	if !memberExists {
		http.Error(w, "Member not found in the project", http.StatusBadRequest)
		return
	}

	// Uklonite člana
	update := bson.M{"$pull": bson.M{"members": req.MemberID}, "$set": bson.M{"updated_at": time.Now()}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": projectObjectID}, update)
	if err != nil {
		http.Error(w, "Failed to remove member", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Member removed successfully")
}

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	// Povezivanje na kolekciju "projects"
	collection := db.Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Preuzimanje svih projekata
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var projects []Project
	if err = cursor.All(ctx, &projects); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Slanje odgovora kao JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
