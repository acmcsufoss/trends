package discord

import (
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/ethanthatonekid/trends/preprocess"
)

// ChannelID is a Discord channel ID.
type ChannelID = discord.ChannelID

// client is a Discord client.
type client struct {
	*api.Client
}

// New creates a new Discord client.
func New(token string) preprocess.ChatProvider {
	return &client{api.NewClient(token)}
}

// ValidateChannels validates the channels.
func (c *client) ValidateChannels(ids []string) error {
	return nil
}

func (c *client) Channel(id string) (preprocess.Channel, error) {
	channelID, err := discord.ParseSnowflake(id)
	if err != nil {
		return preprocess.Channel{}, err
	}

	channel, err := c.Client.Channel(discord.ChannelID(channelID))
	if err != nil {
		return preprocess.Channel{}, err
	}

	return preprocess.Channel{
		ID:   channel.ID.String(),
		Name: channel.Name,
	}, nil
}

// Messages reads a range of messages from the channel.
// The range is inclusive of the start and end timestamps.
func (c *client) Messages(id string, start, end time.Time) (preprocess.MessageList, error) {
	channelID, err := discord.ParseSnowflake(id)
	if err != nil {
		return preprocess.MessageList{}, err
	}

	fromID := discord.MessageID(discord.NewSnowflake(start))
	toID := discord.MessageID(discord.NewSnowflake(end))

	messages, err := c.messagesBefore(discord.ChannelID(channelID), fromID, toID)
	if err != nil {
		return preprocess.MessageList{}, err
	}

	authors := make(map[string]preprocess.Author, len(messages))
	for _, message := range messages {
		authors[message.Author.ID.String()] = convertAuthor(message.Author)
		for _, mentioned := range message.Mentions {
			authors[mentioned.ID.String()] = convertAuthor(mentioned.User)
		}
	}

	return preprocess.MessageList{
		Messages: convertList(messages, convertMessage),
		Authors:  authors,
	}, nil
}

func (c *client) messagesBefore(channelID discord.ChannelID, fromID, toID discord.MessageID) ([]discord.Message, error) {
	messages := make([]discord.Message, 0, 100)

searchLoop:
	for {
		// Search from the latest message to the oldest message.
		// This is because the API returns messages in reverse order.
		messagesBuffer, err := c.MessagesBefore(channelID, toID, 100)
		if err != nil {
			return nil, err
		}

		if len(messagesBuffer) == 0 {
			break
		}

		for i := range messagesBuffer {
			// If the message is older than the start timestamp, we're done.
			if messagesBuffer[i].ID < fromID {
				break searchLoop
			}

			messages = append(messages, messagesBuffer[i])
		}

		// Grab the latest ID in our buffer. We'll fetch messages that are
		// later than this one.
		toID = messages[len(messages)-1].ID
	}

	return messages, nil
}

func convertMessage(message discord.Message) preprocess.Message {
	return preprocess.Message{
		Timestamp: message.Timestamp.Time(),
		AuthorID:  message.Author.ID.String(),
		Text:      message.Content,
	}
}

func convertAuthor(author discord.User) preprocess.Author {
	return preprocess.Author{
		Name:   author.Username,
		ID:     author.ID.String(),
		Avatar: author.AvatarURL(),
	}
}

func convertList[T1 any, T2 any](list []T1, f func(T1) T2) []T2 {
	var newList []T2
	for _, item := range list {
		newList = append(newList, f(item))
	}
	return newList
}
