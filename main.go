package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.HTTPErrorHandler = myErrorHandler

	for _, filename := range configs {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		var call_slack CallSlack
		if err := json.Unmarshal(bytes, &call_slack); err != nil {
			log.Fatal(err)
		}
		e.GET("/"+call_slack.Endpoint, call_slack.Post)

		e.Start(":8000")
	}
}

func myErrorHandler(err error, context echo.Context) {
	//code := http.StatusInternalServerError
	context.Logger().Error(err)
}
