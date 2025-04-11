package models

type AuthUser struct {
	Email    string `json:"email"    binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}
