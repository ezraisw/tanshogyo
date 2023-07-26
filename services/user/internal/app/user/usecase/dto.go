package usecase

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Field struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type AuthenticationResult struct {
	Token string `json:"token"`
}
