package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"

	"tantalic.com/meetup-reminder/internal/meetup"
	"tantalic.com/meetup-reminder/internal/slack"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "meetup.com event reminder"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "slack-webhook",
			EnvVar: "SLACK_WEBHOOK",
			Usage:  "url for slack webhook",
		},
		cli.StringFlag{
			Name:   "meetup",
			EnvVar: "MEETUP_NAME",
			Usage:  "url name of the meetup.com group",
		},
		cli.IntFlag{
			Name:   "days-before-reminder",
			EnvVar: "DAYS_BEFORE_REMINDER",
			Usage:  "number of days prior to the event for reminder",
			Value:  7,
		},
		cli.StringFlag{
			Name:   "message",
			EnvVar: "MESSAGE",
			Usage:  "custom message text for reminder",
		},
	}

	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	m := meetup.Client{
		GroupURLName: c.String("meetup"),
	}
	events, err := m.FetchEvents()
	if err != nil {
		fmt.Println("Error fetching meetup.com events: %s", err.Error())
		return err
	}

	for _, event := range events {
		if isInDays(event, c.Int("days-before-reminder")) {
			notify(event, c)
		}
	}

	if err != nil {
		fmt.Println("Error posting to slack: %s", err.Error())
		return err
	}

	return nil
}

func isInDays(event meetup.Event, days int) bool {
	thisYear, thisMonth, thisDay := time.Now().Add(time.Hour * time.Duration(days) * 24).Date()
	eventYear, eventMonth, eventDay := event.Time().Date()

	return thisYear == eventYear &&
		thisMonth == eventMonth &&
		thisDay == eventDay
}

func notify(event meetup.Event, c *cli.Context) error {
	webhook := slack.Webhook{
		URL: c.String("slack-webhook"),
	}

	message := c.String("message")
	if message == "" {
		message = fmt.Sprintf("is in %d days", c.Int("days-before-reminder"))
	}
	date := event.Time().Format("Monday, January 2\n3:04 PM")
	location := fmt.Sprintf("%s\n%s", event.Venue.Name, event.Venue.AddressLine1)

	return webhook.Post(slack.Message{
		Text: fmt.Sprintf("@channel *Reminder: %s %s* :loudspeaker: ", event.Name, message),
		Attachments: []slack.Attachment{
			slack.Attachment{
				Fallback: date,
				Color:    "#3383b8",
				Fields: []slack.Field{
					slack.Field{
						Title: "When",
						Value: date,
						Short: true,
					},
					slack.Field{
						Title: "Where",
						Value: location,
						Short: true,
					},
				},
			},
		},
	})
}
