package structs

// Channel is a channel of text messages.
type Channel struct {
	// ID is the unique identifier of the channel.
	ID string `json:"id"`
	// Name is the name of the channel.
	Name string `json:"name"`
}
