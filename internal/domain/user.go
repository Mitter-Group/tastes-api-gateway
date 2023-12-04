package domain

import "time"

// User represents a user in the domain logic.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPayload struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
	UserType string `json:"user_type"`
}
