package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	//"net/http"
	//"sync"

	//"github.com/ant0ine/go-json-rest/rest"
	"github.com/ashwanthkumar/slack-go-webhook"
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

func (i CallSlack) Post() []error {
	field := slack.Field{Title: i.Title, Value: i.Message}

	attachment := slack.Attachment{}
	attachment.AddField(field)
	color := i.Color
	attachment.Color = &color
	payload := slack.Payload{
		Username:    i.Username,
		Channel:     i.Channel,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(i.Webhook_url, "", payload)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	bytes, err := ioutil.ReadFile("configs/.conf.json")
	if err != nil {
		log.Fatal(err)
	}
	var call_slack CallSlack
	if err := json.Unmarshal(bytes, &call_slack); err != nil {
		log.Fatal(err)
	}

	var er []error
	er = call_slack.Post()
	if er != nil {
		log.Fatal(er)
	}

	//api := rest.NewApi()
	//api.Use(rest.DefaultDevStack...)
}
