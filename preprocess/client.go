package preprocess

import (
	"github.com/ethanthatonekid/trends/preprocess/channel"
	"github.com/ethanthatonekid/trends/preprocess/discordclient"
)

// Client is an interface for fetching messages from a messaging service.
type Client interface {
	// Channel returns a channel by its ID.
	Channel(id string) (channel.Channel, error)
}

// New creates a new Discord preprocessor client.
func NewDiscordClient(token string) Client {
	client := discordclient.New(token)
	return &client
}
