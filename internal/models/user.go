package models

import "time"

// User represents the API response model
type User struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
	Age  int       `json:"age,omitempty"`
}
