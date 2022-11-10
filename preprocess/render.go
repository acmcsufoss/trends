package preprocess

import (
	"time"

	"github.com/pkg/errors"
)

// RenderedChannel is a channel's messages rendered as JSON.
type RenderedChannel struct {
	ChannelID   string                    `json:"channel_id"`
	ChannelName string                    `json:"channel_name"`
	Authors     map[string]RenderedAuthor `json:"authors"`
	Messages    []RenderedMessage         `json:"messages"`
}

type RenderedMessage = Message

type RenderedAuthor struct {
	// Name is the name of the user.
	Name string `json:"name"`
	// Avatar is the URL of the user's avatar.
	Avatar string `json:"avatar"`
}

// RenderChannels renders the channels to JSON.
func RenderChannels(provider ChatProvider, ids []string, start, end time.Time) ([]*RenderedChannel, error) {
	if err := provider.ValidateChannels(ids); err != nil {
		return nil, errors.Wrap(err, "failed to validate channels")
	}

	channels := make([]*RenderedChannel, 0, len(ids))

	for _, id := range ids {
		result, err := RenderChannel(provider, id, start, end)
		if err != nil {
			return nil, err
		}

		channels = append(channels, result)
	}

	return channels, nil
}

// RenderChannel renders a channel to JSON.
func RenderChannel(provider ChatProvider, id string, start, end time.Time) (*RenderedChannel, error) {
	channel, err := provider.Channel(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get channel")
	}

	messages, err := provider.Messages(id, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read messages")
	}

	renderedMsgs := make([]RenderedMessage, len(messages.Messages))
	renderedAuthors := make(map[string]RenderedAuthor)

	for i, msg := range messages.Messages {
		renderedMsgs[i] = RenderedMessage{
			Timestamp: msg.Timestamp,
			AuthorID:  msg.AuthorID,
			Text:      msg.Text,
		}

		renderedAuthors[msg.AuthorID] = RenderedAuthor{
			Name:   messages.Authors[msg.AuthorID].Name,
			Avatar: messages.Authors[msg.AuthorID].Avatar,
		}
	}

	return &RenderedChannel{
		ChannelID:   id,
		ChannelName: channel.Name,
		Authors:     renderedAuthors,
		Messages:    renderedMsgs,
	}, nil
}
