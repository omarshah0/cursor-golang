package types

// Book represents a book in the system
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    // Add other fields as necessary
}