package models

// User or Subscriber who joins a news letter.
type User struct {
	Id string

	// Email that gets used when subscribing to an email.
	Email string

	// all the newsletters that user subscribes to in the Gabriel network.
	Newsletters []string
}
