package main

import (
	"fmt"
	"net/http"

	"cyoa/jsonparser"

	"cyoa/template"
)

func main() {

	storyMap := jsonparser.Unmarshal("./jsonparser/gopher.json")

	for k, v := range storyMap {
        h := template.AdventureHandler(v)
        http.HandleFunc("/"+k, h)
	}

    err := http.ListenAndServe(":8881", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
