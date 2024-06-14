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

type ProviderData struct {
	Provider       string `json:"provider" dynamodbav:"Provider"`
	ProviderUserID string `json:"provider_user_id" dynamodbav:"ProviderUserID"`
	UserFullname   string `json:"user_fullname" dynamodbav:"UserFullname"`
	Email          string `json:"email" dynamodbav:"Email"`
}

type UserPayload struct {
	ID             string         `json:"ID" dynamodbav:"ID"`
	Email          string         `json:"email" dynamodbav:"Email"`
	UserFullname   string         `json:"user_fullname" dynamodbav:"UserFullname"`
	ProfilePicture string         `json:"profile_picture" dynamodbav:"ProfilePicture"`
	Providers      []ProviderData `json:"providers" dynamodbav:"Providers"`
}
