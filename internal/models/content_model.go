package models

type ContentType int

const (
	Text ContentType = iota
	Markdown
	Html
)

// Email that gets sent to every user in a given newsletter.
type Content struct {
	Id string `json:"Id"`

	// Display title, or Email subject.
	Title string `json:"Title"`

	// Newsletter list that this content is attached to.
	NewsletterId []string `json:"NewsletterId"`

	// Unix timestamp for when the email should be sent to the users.
	ReleaseDate int64 `json:"ReleaseDate"`

	// Content type (.txt, .md, or .html file).
	Type ContentType `json:"Type"`

	// Content ID. Used for finding the content file in the db.
	ContentId string `json:"ContentId"`

	// Has the content been sent to the user yet.
	SentState bool `json:"SentState"`
}
