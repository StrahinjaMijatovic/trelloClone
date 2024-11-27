package main

import (
	"github.com/go-playground/validator/v10"
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

var validate = validator.New()

type RegisterRequest struct {
	FirstName       string `json:"firstName" validate:"required,alpha"`
	LastName        string `json:"lastName" validate:"required,alpha"`
	Username        string `json:"username" validate:"required,alphanum,min=3,max=20"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Age             int    `json:"age" validate:"required,min=18"`
	Country         string `json:"country" validate:"required"`
	Role            string `json:"role" validate:"required,oneof=NK M C"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
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
