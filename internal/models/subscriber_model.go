package models

// Subscriber who joins a newsletter.
type Subscriber struct {
	// UUID for the unique user.
	UUId string

	// Email that gets used when subscribing to an email.
	Email string

	// All the newsletters that the user subscribes to in the Gabriel network.
	NewsletterUUIds []string
}
