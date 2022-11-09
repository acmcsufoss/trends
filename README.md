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
# Run the command to see a help menu.
go run .
```

```bash
# Run a sample command.
go run . preprocess --channels "710225253737037836" --output "results" --start "2022-11-07T0:00:00Z" --end "2023-01-01T0:00:00Z"
```

### Build

```bash
# Build the bot
go build
```