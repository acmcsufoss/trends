package main

import (
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/ethanthatonekid/trends/preprocess"
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
	client preprocess.Client
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
					&cli.TimestampFlag{
						Name:    "start",
						Aliases: []string{"s"},
						Usage:   "start date of the range to analyze",
					},
					&cli.TimestampFlag{
						Name:    "end",
						Aliases: []string{"e"},
						Usage:   "end date of the range to analyze",
					},
					&cli.StringSliceFlag{
						Name:    "channels",
						Aliases: []string{"c"},
						Usage:   "the channels to analyze",
					},
				},
				Action: func(ctx *cli.Context) error {
					for _, channel := range ctx.StringSlice("channels") {
						_, err := discord.ParseSnowflake(channel)
						if err != nil {
							return errors.Wrap(err, "failed to parse channel ID")
						}
					}

					app.client = preprocess.NewDiscordClient(os.Getenv("DISCORD_TOKEN"))

					for _, channelID := range ctx.StringSlice("channels") {
						channel, err := app.client.Channel(channelID)
						if err != nil {
							return errors.Wrap(err, "failed to get channel")
						}

						messages, err := channel.Read(ctx.Timestamp("start"), ctx.Timestamp("end"))
						if err != nil {
							return errors.Wrap(err, "failed to read messages")
						}

						for _, message := range messages {
							log.Println(message)
						}
					}

					return nil
				},
			},
		},
	}

	return app
}
