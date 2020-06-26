package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	//"net/http"
	//"sync"
	//"github.com/ant0ine/go-json-rest/rest"
)

func getConfigs(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			path, err := getConfigs(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, path...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, nil
}

func main() {
	configs, err := getConfigs("./configs/")
	fmt.Println(configs)
	if err != nil {
		log.Fatal(err)
		return
	}

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
