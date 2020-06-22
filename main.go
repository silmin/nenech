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
	Name    string `json:"name"`
	Channel string `json:"channel"`
	Message string `json:"message"`
}

func (i CallSlack) Post() []error {
	field := slack.Field{Title: "Message", Value: i.Message}

	attachment := slack.Attachment{}
	attachment.AddField(field)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Username:    USERNAME,
		Channel:     i.Channel,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(webhook_urls[i.Channel], "", payload)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	bytes, err := ioutil.ReadFile("configs/test.json")
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
