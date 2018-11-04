package scalyr

import (
	"encoding/json"
	"github.com/convertkit/stories/stories"
  "bytes"
)

type Event stories.Story

func (e *Event) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	data["ts"] = e.Timestamp
	data["sev"] = e.Severity

  attributes := e.Data
  attributes["message"] = jsonValidMessage([]byte(e.Message))

  data["attrs"] = attributes

	return json.Marshal(data)
}

func jsonValidMessage(message []byte) []byte {
  s := [][]byte{
    []byte(`"`),
    message,
    []byte(`"`)}

  return bytes.Join(s, []byte(""))
}
