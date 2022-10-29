package models

type Register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserRole string `json:"user_role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
