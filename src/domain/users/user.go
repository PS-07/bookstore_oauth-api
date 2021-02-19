package users

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// UserLoginRequest struct
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
