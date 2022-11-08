package discordclient

import (
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/ethanthatonekid/trends/preprocess/channel"
	"github.com/ethanthatonekid/trends/preprocess/structs"
)

// ChannelID is a Discord channel ID.
type ChannelID = discord.ChannelID

// client is a Discord client.
type client struct {
	*api.Client
}

// discordChannel implements a Discord Channel.
type discordChannel struct {
	channel.Channel
	client  client
	channel *discord.Channel
}

// New creates a new Discord client.
func New(token string) client {
	return client{api.NewClient(token)}
}

// Channel returns a channel by its ID.
func (c *client) Channel(id string) (channel.Channel, error) {
	chID, err := discord.ParseSnowflake(id)
	if err != nil {
		return nil, err
	}

	ch, err := c.Client.Channel(discord.ChannelID(chID))
	if err != nil {
		return nil, err
	}

	return newDiscordChannel(ch, *c), nil
}

func newDiscordChannel(ch *discord.Channel, client client) channel.Channel {
	return &discordChannel{channel: ch, client: client}
}

// Read reads a range of messages from the channel.
// The range is inclusive of the start and end timestamps.
func (c *discordChannel) Read(start, end *time.Time) ([]structs.Message, error) {
	messages, err := c.findMessages(func(msg *discord.Message) bool {
		return start.After(msg.Timestamp.Time()) && end.Before(msg.Timestamp.Time())
	})
	if err != nil {
		return nil, err
	}

	return convertMessages(messages), nil
}

// findMessages returns the first consecutive set of messages that match the filter.
func (c *discordChannel) findMessages(filter func(msg *discord.Message) bool) ([]*discord.Message, error) {
	var messages []*discord.Message
	var lastID discord.MessageID
	var err error

	messagesBuffer := make([]discord.Message, 0, 100)
	for {
		messagesBuffer, err = c.client.MessagesAfter(c.channel.ID, lastID, 100)
		if err != nil {
			return nil, err
		}

		if len(messagesBuffer) == 0 {
			break
		}

		for i := range messagesBuffer {
			message := messagesBuffer[i]
			if filter(&message) {
				messages = append(messages, &message)
			} else if len(messages) > 0 {
				break
			}
		}

		lastID = messagesBuffer[len(messagesBuffer)-1].ID
	}

	return messages, nil
}

// convertMessages converts a slice of Discord messages to a slice of
// internal messages.
func convertMessages(messages []*discord.Message) []structs.Message {
	var internalMessages []structs.Message

	for _, message := range messages {
		internalMessage := structs.Message{
			Timestamp: message.Timestamp.Time(),
			Author: structs.Author{
				Name:   message.Author.Username,
				ID:     message.Author.ID.String(),
				Avatar: message.Author.AvatarURL(),
			},
			Text: message.Content,
		}
		internalMessages = append(internalMessages, internalMessage)
	}

	return internalMessages
}
