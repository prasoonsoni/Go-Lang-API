package main

// Model for course - file
type Note struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Fake DB
var notes []Note

// middleware, helper - file
func (n Note) IsEmpty() bool {
	return n.Title == "" && n.Description == ""
}

func main() {

}
