package schemas

import "github.com/google/uuid"

// SignUpRequest is the request body for sign up
type SignUpRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

// SignUpResponse is the response body for sign up
type SignUpResponse struct {
	UserID uuid.UUID `json:"userID"`
}

// SignInRequest is the request body for sign in
type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignInResponse is the response body for sign in
type SignInResponse struct {
	Token string `json:"token"`
}
