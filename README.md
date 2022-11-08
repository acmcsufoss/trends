# trends
Produce trends analysis based on messages shared on Discord.

## Dev

### Setup

```bash
# Install dependencies
go mod tidy
```

### Run

Set up a `.env` file with the following variables:

```bash
# Discord bot token
DISCORD_TOKEN=your_token_here
```

Our development bot invitation link: <https://discord.com/api/oauth2/authorize?client_id=1039347363849977867&permissions=1024&scope=bot>.

```bash
# Run the bot
go run main.go
```

### Build

```bash
# Build the bot
go build -o trends
```