package models

import "time"

type User struct {
	TUID           uint32    `json:"tuid"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashedPassword"`
	FirstName      string    `json:"firstName"`
	MiddleName     string    `json:"middleName"`
	LastName       string    `json:"lastName"`
	Departments    []string  `json:"departments"`
	IsAdmin        bool      `json:"isAdmin"`
	IsCoordinator  bool      `json:"isCoordinator"`
	IsFaculty      bool      `json:"isFaculty"`
	IsStaff        bool      `json:"isStaff"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
