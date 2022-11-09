package preprocess

import (
	"time"
)

// ChatProvider provides a channel by ID from any service.
type ChatProvider interface {
	// ValidateChannels validates the channels.
	ValidateChannels(ids []string) error
	// Channel gets a channel by ID.
	Channel(id string) (Channel, error)
	// Messages reads a range of messages from the channel.
	Messages(id string, start, end time.Time) (MessageList, error)
}

type Channel struct {
	ID   string
	Name string
}

type MessageList struct {
	Messages []Message
	Authors  map[string]Author
}

// Message is a text message struct.
type Message struct {
	// Timestamp is the time the message was sent.
	Timestamp time.Time `json:"timestamp"`
	// AuthorID is the user who sent the message.
	AuthorID string `json:"author_id"`
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
