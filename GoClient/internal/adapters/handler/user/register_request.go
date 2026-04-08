package user

// Define structures for clien's JSON
type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Nickname  string `json:"nickname"   binding:"required"`
	Email     string `json:"email"      binding:"required,email"`
	Password  string `json:"password"   binding:"required,min=8"`
	DateBirth string `json:"date_birth" binding:"required,datetime=2006-01-02"`
	Gender    string `json:"gender"     binding:"required,oneof=Male Female Other"`
}
