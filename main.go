package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/silmin/nenech/handler"
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
	if err != nil {
		log.Fatal(err)
		return
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = handler.MyErrorHandler

	for _, filename := range configs {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		var call_slack CallSlack
		if err := json.Unmarshal(bytes, &call_slack); err != nil {
			log.Fatal(err)
		}
		ep_tmp := "/%s"
		endpoint := fmt.Sprintf(ep_tmp, call_slack.Endpoint)
		e.GET(endpoint, call_slack.Post)

	}
	port := flag.String("port", "80", "starting port")
	flag.Parse()

	e.Start(":" + *port)
}
