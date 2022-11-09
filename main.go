package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/ethanthatonekid/trends/preprocess"
	"github.com/ethanthatonekid/trends/preprocess/discord"
)

func main() {
	godotenv.Load(".env")
	app := NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type App struct {
	*cli.App
}

func NewApp() *App {
	app := &App{}

	app.App = &cli.App{
		Name:     "trends",
		HelpName: "discover trends in your Discord server",
		Commands: []*cli.Command{
			{
				Name:     "preprocess",
				HelpName: "preprocesses messages for analysis",
				Usage:    "trends preprocess <channel ID> <start timestamp> <end timestamp>",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "channels",
						Aliases: []string{"c"},
						Usage:   "the channels to analyze",
					},
					&cli.PathFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "the output directory",
						Value:   "results",
					},
					&cli.TimestampFlag{
						Name:    "start",
						Aliases: []string{"s"},
						Usage:   "start date of the range to analyze (see https://unixtimestamp.com/)",
						Layout:  time.RFC3339,
					},
					&cli.TimestampFlag{
						Name:    "end",
						Aliases: []string{"e"},
						Usage:   "end date of the range to analyze (see https://unixtimestamp.com/)",
						Layout:  time.RFC3339,
					},
				},
				Action: preprocessAction,
			},
		},
	}

	return app
}

func preprocessAction(ctx *cli.Context) error {
	provider := discord.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	channelIDs := ctx.StringSlice("channels")
	outputRoot := ctx.Path("output")
	start := ctx.Timestamp("start").UTC()
	end := ctx.Timestamp("end").UTC()

	// timeFormat (generated from https://godate.diamondb.xyz/)
	const timeFormat = "06-01-02_15-04-05"
	timeRange := fmt.Sprintf("%s-%s", start.Format(timeFormat), end.Format(timeFormat))

	// Fetch and convert messages from each channel.
	channels, err := preprocess.RenderChannels(provider, channelIDs, start, end)
	if err != nil {
		return errors.Wrap(err, "failed to render messages")
	}

	// Write each channel to a file.
	for _, ch := range channels {
		b, err := json.MarshalIndent(ch, "", "  ")
		if err != nil {
			return errors.Wrap(err, "failed to marshal channel")
		}

		filename := fmt.Sprintf("%s.json", ch.ChannelID)
		pathname := filepath.Join(outputRoot, timeRange, filename)
		if err = os.MkdirAll(filepath.Dir(pathname), 0755); err != nil {
			return errors.Wrap(err, "failed to create directory")
		}

		if err = ioutil.WriteFile(pathname, b, 0644); err != nil {
			return errors.Wrap(err, "failed to write messages to file")
		}
	}

	return nil
}
