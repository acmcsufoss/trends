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
# Discord bot token.
DISCORD_TOKEN=your_token_here
```

Our development bot invitation link: <https://discord.com/api/oauth2/authorize?client_id=1039347363849977867&permissions=1024&scope=bot>.

```bash
# Run the command to see a help menu.
go run .
```

```bash
# Run a sample command (#general).
go run . preprocess --channels "710225253737037836" --output "results" --start "2022-11-01T0:00:00Z" --end "2022-11-02T0:00:00Z"

# Run a sample command with aliases (#resume-check)
go run . preprocess -c "745714911056756849" -o "results" -s "2022-01-01T0:00:00Z" -e "2023-01-01T0:00:00Z"
```

```bash
# Run a sample command to preprocess major channels into months.
# general: 710225253737037836
# off-topic: 999910900121223178
# opportunities: 711075012672487514
# resources: 802312854811443231
# coding-help: 745456552718106654
# resume-check: 745714911056756849
# interview-prep: 745797700582375465
# memes: 711074976840679515
# music-recs: 987785411373973535
# sports: 970417311129415710
# gaming: 1039307766554185728
# student-schedules: 903343163895341066
# server-suggestions: 742567658628841555

# October 2022
go run . preprocess -s "2022-10-01T0:00:00Z" -e "2022-11-01T0:00:00Z" -c "710225253737037836,999910900121223178,711075012672487514,802312854811443231,745456552718106654,745714911056756849,745797700582375465,711074976840679515,987785411373973535,970417311129415710,1039307766554185728,903343163895341066,742567658628841555"
```

### Build

```bash
# Build the bot.
go build
```

### Test

```bash
# Run tests.
go test ./...
```

---

Programmed with ❤️ by <https://acmcsuf.com> **Bot Committee**!
