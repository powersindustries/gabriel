package models

// Email that gets sent to every user in a given newsletter.
type Content struct {
	UUId string

	// Display title, or Email subject.
	Title string

	// Unix timestamp for when the email should be sent to the users.
	ReleaseDate int64

	// Filetype (.txt, .md, or .html file).
	Type string

	// Newsletter UUID that this content is attached to.
	NewsletterUUId string
}
