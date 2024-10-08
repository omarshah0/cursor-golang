package types

// User represents a user in the system
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    // Add other fields as necessary
}