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
	ID             string `json:"ID" dynamodbav:"ID"`
	Provider       string `json:"provider" dynamodbav:"Provider"`
	ProviderUserID string `json:"provider_user_id" dynamodbav:"ProviderUserID"`
	UserFullname   string `json:"user_fullname" dynamodbav:"UserFullname"`
	Email          string `json:"email" dynamodbav:"Email"`
}
