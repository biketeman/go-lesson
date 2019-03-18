package blog

type Post struct {
	Title       string
	Description string
	Content     string
	Authors     []*User
}

type User struct {
	FirstName string
	LastName  string
	Username  string
	Age       int
	Email     string
}

type Blog interface {
	NewPost(*Post) error
	UpdatePost(*Post) error
	DeletePost(*Post) error
	// AddCollaborators()
	// NewUser()
	// UpdateUser()
	// DeleteUser()
}
