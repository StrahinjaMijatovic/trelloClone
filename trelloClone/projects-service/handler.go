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

	project.ID = primitive.NewObjectID()
	project.Members = []string{}
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
		log.Printf("ID projekta ne vazi: %v", err)
		http.Error(w, "ID projekta ne vazi", http.StatusBadRequest)
		return
	}

	var project Project
	if err := collection.FindOne(ctx, bson.M{"_id": projectObjectID}).Decode(&project); err != nil {
		log.Printf("Projekat nije pronadjen: %v", err)
		http.Error(w, "Projekat nije pronadjen", http.StatusNotFound)
		return
	}

	log.Printf("Projekat pronadjen: %+v", project)

	// Validacije
	if project.ManagerID != req.ManagerID {
		log.Printf("Autorizacija menadzera ne uspesna: %s", req.ManagerID)
		http.Error(w, "Autorizacija menadzera ne uspesna", http.StatusForbidden)
		return
	}

	if project.EndDate.Before(time.Now()) {
		log.Printf("Nije moguce dodati clanove na gotov projekat")
		http.Error(w, "Nije moguce dodati clanove na gotov projekat", http.StatusBadRequest)
		return
	}

	if len(project.Members) >= project.MaxMembers {
		log.Printf("Projekat je na maximalnom kapacitetu")
		http.Error(w, "Projekat je na maximalnom kapacitetu", http.StatusBadRequest)
		return
	}

	for _, member := range project.Members {
		if member == req.MemberID {
			log.Printf("Clan je vec dodat: %s", req.MemberID)
			http.Error(w, "Clan je vec dodat", http.StatusBadRequest)
			return
		}
	}

	update := bson.M{"$push": bson.M{"members": req.MemberID}, "$set": bson.M{"updated_at": time.Now()}}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": projectObjectID}, update)
	if err != nil {
		log.Printf("Failed to update project: %v", err)
		http.Error(w, "Failed to add member", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Member added successfully")
}

func RemoveMemberHandler(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["id"]
	var req struct {
		MemberID  string `json:"memberId"`
		ManagerID string `json:"managerId"`
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

	if project.ManagerID != req.ManagerID {
		http.Error(w, "You are not authorized to manage this project", http.StatusForbidden)
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
