package models

// Top level class for the Newsletter.
type Newsletter struct {
	UUId string

	// Display name of the Newsletter.
	Name string

	// Short description of what the newsletter is all about.
	Description string

	// Contact for the newsletter.
	ContactEmail string

	// List of all the subscribers who subscribe to this newsletter by UUId.
	SubscriberList []string
}
