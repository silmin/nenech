package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	//"net/http"
	//"sync"
	//"github.com/ant0ine/go-json-rest/rest"
)

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
