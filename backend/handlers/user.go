package handlers

import "github.com/atbu/slate/backend/models"

type UserHandler struct {
	userRepo *models.UserRepository
}

func NewUserHandler(userRepo *models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// Placeholder for now, will be used for user profile / account settings most likely
