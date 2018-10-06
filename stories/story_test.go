package stories

import (
	"encoding/json"
	"testing"
)

func TestNewStoryWithInvalidJSON(t *testing.T) {
	story, err := NewStory([]byte("Invalid JSON"))

	if story != nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func TestNewStoryWithValidJSON(t *testing.T) {
	story, err := NewStory([]byte("{\"foo\": \"bar\"}"))

	if story == nil {
		t.Error(err)
	}

	if err != nil {
		t.Fail()
	}
}

func TestNewStoryHasDefaultSeverity(t *testing.T) {
	story, err := NewStory([]byte("{\"foo\": \"bar\"}"))

	if err != nil {
		t.Fail()
	}

	if story.Severity != 3 {
		t.Error(story)
	}
}

func TestNewStorySetsSeverityIfExists(t *testing.T) {
	sev := 4
	data := make(map[string]interface{})
	data["severity"] = sev

	json, err := json.Marshal(data)

	if err != nil {
		t.Error(err)
	}

	story, err := NewStory(json)

	if err != nil {
		t.Fail()
	}

	if story.Severity != 4 {
		t.Error(story)
	}
}

func TestNewStoryDataIsNeverNil(t *testing.T) {
	data := make(map[string]interface{})

	json, err := json.Marshal(data)

	if err != nil {
		t.Error(err)
	}

	story, err := NewStory(json)

	if err != nil {
		t.Fail()
	}

	if story.Data == nil {
		t.Fail()
	}
}
