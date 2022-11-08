package channel

import (
	"time"

	"github.com/ethanthatonekid/trends/preprocess/structs"
)

// Channel is a channel of text messages.
// Channel is implemented by a Discord or Slack text channel.
type Channel interface {
	// Read reads a range of messages from the channel.
	Read(start, end *time.Time) ([]structs.Message, error)
}
