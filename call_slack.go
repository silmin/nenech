package main

import (
	"log"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/labstack/echo"
)

type CallSlack struct {
	Endpoint    string `json:"endpoint"`
	Username    string `json:"username"`
	Channel     string `json:"channel"`
	Webhook_url string `json:"webhook_url"`
	Title       string `json:"title"`
	Message     string `json:"message"`
	Color       string `json:"color"`
}

func (i CallSlack) Post(context echo.Context) error {
	field := slack.Field{Title: i.Title, Value: i.Message}

	attachment := slack.Attachment{}
	attachment.AddField(field)
	attachment.Color = &i.Color
	payload := slack.Payload{
		Username:    i.Username,
		Channel:     i.Channel,
		Attachments: []slack.Attachment{attachment},
	}

	errs := slack.Send(i.Webhook_url, "", payload)
	if errs != nil {
		for _, err := range errs {
			log.Output(1, err.Error())
		}
		return errs[0]
	}

	return nil
}
