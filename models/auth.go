package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// Request
type SignUpParams struct {
	TUID          uint32   `json:"tuid"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	FirstName     string   `json:"firstName"`
	MiddleName    string   `json:"middleName"`
	LastName      string   `json:"lastName"`
	Departments   []string `json:"departments"`
	IsAdmin       bool     `json:"isAdmin"`
	IsCoordinator bool     `json:"isCoordinator"`
	IsFaculty     bool     `json:"isFaculty"`
	IsStaff       bool     `json:"isStaff"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWT claims
type AccountClaims struct {
	TUID uint32 `json:"tuid"`
	jwt.RegisteredClaims
}
