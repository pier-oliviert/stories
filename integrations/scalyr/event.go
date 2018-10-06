package scalyr

import (
	"encoding/json"
	"github.com/pothibo/stories/stories"
)

type Event stories.Story

func (e *Event) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	data["ts"] = e.Timestamp
	data["sev"] = e.Severity
	data["attrs"] = e.Data

	return json.Marshal(data)
}
