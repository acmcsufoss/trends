package structs

import (
	"time"
)

// Message is a text message struct.
type Message struct {
	// Timestamp is the time the message was sent.
	Timestamp time.Time `json:"timestamp"`
	// Author is the user who sent the message.
	Author Author `json:"author"`
	// Text is the content of the message.
	Text string `json:"text"`
}

// Author is a user struct.
type Author struct {
	// Name is the name of the user.
	Name string `json:"name"`
	// ID is the unique identifier of the user.
	ID string `json:"id"`
	// Avatar is the URL of the user's avatar.
	Avatar string `json:"avatar"`
}
