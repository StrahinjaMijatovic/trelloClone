package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role string

const (
	Unauthenticated Role = "NK"
	Host            Role = "M"
	Guest           Role = "C"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName"`
	Username     string             `json:"username" bson:"username"`
	PasswordHash string             `json:"passwordHash" bson:"passwordHash"`
	Email        string             `json:"email" bson:"email"`
	Age          int                `json:"age" bson:"age"`
	Country      string             `json:"country" bson:"country"`
	Role         Role               `json:"role" bson:"role"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type RegisterRequest struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
	Age             int    `json:"age"`
	Country         string `json:"country"`
	Role            string `json:"role"` // NK, M, C
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type UpdateProfileRequest struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Username        string `json:"username"`
	OldPassword     string `json:"oldPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	Email           string `json:"email"`
	Age             int    `json:"age"`
	Country         string `json:"country"`
}
