package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Story struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}

func Unmarshal(path string) map[string]Story {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var s map[string]Story

	if err := json.Unmarshal(bytes, &s); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Json Unmarshal Failure", err)
	}
	return s
}
