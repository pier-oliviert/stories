package main

import "encoding/json"

type Story struct {
	Severity  int
	Message   string
	Timestamp string
	Data      map[string]string
}

func NewStoryFromJSON(bytes []byte) (*Story, error) {
	var story Story
	err := json.Unmarshal(bytes, &story)

	if err != nil {
		return nil, err
	}

	return &story, nil
}
