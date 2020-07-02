package main

import (
	"log"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/labstack/echo"
)

type (
	CallSlack struct {
		Endpoint    string `json:"endpoint"`
		Channel     string `json:"channel"`
		Webhook_url string `json:"webhook_url"`

		Username string `json:"username"`
		Title    string `json:"title"`
		Message  string `json:"message"`
		Color    string `json:"color"`
	}

	CustomBinder struct{}
)

func (cb *CustomBinder) Bind(i *CallSlack, context echo.Context) error {
	if param := context.QueryParam("username"); param != "" {
		i.Username = param
	}
	if param := context.QueryParam("title"); param != "" {
		i.Title = param
	}
	if param := context.QueryParam("message"); param != "" {
		i.Message = param
	}
	if param := context.QueryParam("color"); param != "" {
		i.Color = param
	}
	if param := context.FormValue("username"); param != "" {
		i.Username = param
	}
	if param := context.FormValue("title"); param != "" {
		i.Title = param
	}
	if param := context.FormValue("message"); param != "" {
		i.Message = param
	}
	if param := context.FormValue("color"); param != "" {
		i.Color = param
	}
	return nil
}

func (i CallSlack) Post(context echo.Context) error {
	cb := new(CustomBinder)
	if err := cb.Bind(&i, context); err != nil {
		return err
	}

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
