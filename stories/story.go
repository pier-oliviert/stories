package stories

import (
	"encoding/json"
)

type Story struct {
	Severity  int
	Message   string
	Timestamp string
	Data      map[string]string
}

func NewStory(bytes []byte) (*Story, error) {
	var story Story
	err := json.Unmarshal(bytes, &story)

	if err != nil {
		return nil, err
	}

	if story.Data == nil {
		story.Data = make(map[string]string)
	}

	if story.Severity == 0 {
		story.Severity = 3
	}

	return &story, nil
}
