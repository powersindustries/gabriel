package models

// Top level class for the Newsletter.
type Newsletter struct {
	Id string

	// Display name of the Newsletter.
	Name string

	// Short description of what the newsletter is all about.
	Description string

	// Contact for the newsletter.
	ContactEmail string

	// Subscriber list.
	UserList []string
}
