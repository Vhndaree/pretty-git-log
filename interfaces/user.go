package interfaces

// User - struct for User Info
type User struct {
	Name     string `json:"name"`
	UserName string `json:"login"`
	Email    string `json:"email"`
}
